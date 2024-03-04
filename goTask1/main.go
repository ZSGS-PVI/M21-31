package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/hpcloud/tail"
)

type NginxLogEntry struct {
	LogEntry LogEntry `json:"nginx_logs"`
}

type LogEntry struct {
	RemoteAddr       string `json:"CLIENT IP"`
	RemoteUser       string `json:"REMOTE USER"`
	TimeLocal        string `json:"DATE & TIME"`
	Request          string `json:"REQUEST"`
	Status           int    `json:"TARGET STATUS CODE"`
	BodyBytesSent    int    `json:"BODY BYTES SENT"`
	Referer          string `json:"REFERER"`
	UserAgent        string `json:"USER AGENT"`
	SslProtocol      string `json:"SSL PROTOCOL"`
	SslCipher        string `json:"SSL CIPHER"`
	RequestTime      string `json:"REQUEST PROCESSING TIME"`
	UpstreamRespTime string `json:"TARGET PROCESSING TIME"`
	RequestLength    int    `json:"RECEIVED BYTES"`
	UpstreamAddr     string `json:"TARGET IP:PORT"`
	UpstreamStatus   int    `json:"STATUS CODE FROM LOAD BALANCER"`
	BytesSent        int    `json:"BYTES SENT"`
	ServerName       string `json:"SNI DOMAIN"`
	Host             string `json:"HOST"`
	RequestURI       string `json:"URI"`
	HttpUpgrade      string `json:"HTTP UPGRADE"`
}

var (
	lastOffset          int64
	processedLogEntries = make(map[string]struct{})
	mutex               sync.Mutex
)

func main() {
	logFilePath := "/var/log/nginx/access.log" // Adjust as needed

	// WebSocket connection to Java servlet
	conn, _, err := websocket.DefaultDialer.Dial("ws://vishnu:8087/RedisWebSoc/serverws", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	// Ticker for reading log file at regular intervals 
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// Continuously read the Nginx access log file and send the parsed log entries over WebSocket
	for range ticker.C {
		processLogFile(logFilePath, conn)
	}
}

// processLogFile reads the Nginx access log file, parses log entries, and sends them over WebSocket
func processLogFile(logFilePath string, conn *websocket.Conn) {
	// Open the Nginx access log file
	tailConfig := tail.Config{Location: &tail.SeekInfo{Offset: lastOffset, Whence: 0}, ReOpen: true, Follow: true}
	t, err := tail.TailFile(logFilePath, tailConfig)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer t.Stop()

	for line := range t.Lines {
		// Parse each log entry into JSON format
		logEntry, err := parseLogEntry(line.Text)
		if err != nil {
			log.Printf("Error parsing log entry: %v", err)
			continue
		}

		// Check if log entry has already been processed
		entryID := getLogEntryID(logEntry)
		if isProcessed(entryID) {
			continue
		}

		// Convert log entry to JSON with "nginx_logs" key
		jsonData, err := json.Marshal(map[string]interface{}{"nginx_logs": logEntry})
		if err != nil {
			log.Printf("Error marshalling JSON: %v", err)
			continue

		}
		fmt.Println(logEntry)
		// Send the JSON data over WebSocket connection
		err = conn.WriteMessage(websocket.TextMessage, jsonData)
		if err != nil {
			log.Println("write:", err)
			break
		}

		// Mark log entry as processed
		markProcessed(entryID)
	}
}

// parseLogEntry parses a single log entry from Nginx log format into a LogEntry struct
func parseLogEntry(line string) (LogEntry, error) {
	parts := strings.Split(line, "~")
	if len(parts) != 20 {
		return LogEntry{}, nil // Skip invalid lines
	}

	status, _ := strconv.Atoi(parts[4])
	bodyBytesSent, _ := strconv.Atoi(parts[5])
	requestLength, _ := strconv.Atoi(parts[12])
	upstreamStatus, _ := strconv.Atoi(parts[14])
	bytesSent, _ := strconv.Atoi(parts[15])

	logEntry := LogEntry{
		RemoteAddr:       parts[0],
		RemoteUser:       parts[1],
		TimeLocal:        parts[2],
		Request:          parts[3],
		Status:           status,
		BodyBytesSent:    bodyBytesSent,
		Referer:          parts[6],
		UserAgent:        parts[7],
		SslProtocol:      parts[8],
		SslCipher:        parts[9],
		RequestTime:      parts[10],
		UpstreamRespTime: parts[11],
		RequestLength:    requestLength,
		UpstreamAddr:     parts[13],
		UpstreamStatus:   upstreamStatus,
		BytesSent:        bytesSent,
		ServerName:       parts[16],
		Host:             parts[17],
		RequestURI:       parts[18],
		HttpUpgrade:      parts[19],
	}

	return logEntry, nil
}

// getLogEntryID generates a unique identifier for a log entry
func getLogEntryID(logEntry LogEntry) string {
	// Customize this function to generate a unique identifier based on log entry fields
	return logEntry.RemoteAddr + "-" + logEntry.TimeLocal + "-" + logEntry.RequestURI
}

// isProcessed checks if a log entry has already been processed
func isProcessed(entryID string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	_, ok := processedLogEntries[entryID]
	return ok
}

// markProcessed marks a log entry as processed
func markProcessed(entryID string) {
	mutex.Lock()
	defer mutex.Unlock()
	processedLogEntries[entryID] = struct{}{}
}
