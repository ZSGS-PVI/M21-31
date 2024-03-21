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
	RemoteAddr       string `json:"client_ip"`
	RemoteUser       string `json:"remote_user"`
	TimeLocal        string `json:"date_time"`
	Request          string `json:"request"`
	Status           int    `json:"target_status_code"`
	BodyBytesSent    int    `json:"body_bytes_sent"`
	Referer          string `json:"referer"`
	UserAgent        string `json:"user_agent"`
	SslProtocol      string `json:"ssl_protocol"`
	SslCipher        string `json:"ssl_cipher"`
	RequestTime      string `json:"request_processingtime"`
	UpstreamRespTime string `json:"target_processingtime"`
	RequestLength    int    `json:"received_bytes"`
	UpstreamAddr     string `json:"target_ip_port"`
	UpstreamStatus   int    `json:"statuscode_from_loadbalancer"`
	BytesSent        int    `json:"bytes_sent"`
	ServerName       string `json:"sni_domain"`
	Host             string `json:"host"`
	RequestURI       string `json:"uri"`
	HttpUpgrade      string `json:"http_upgrade"`
}

var (
	lastOffset    int64
	mutex         sync.Mutex
	totalLogsSent int
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

	// Print the total number of logs sent
	//fmt.Println("Total logs sent:", totalLogsSent)
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

		sendLogEntry(logEntry, conn)
		totalLogsSent++
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

// sendLogEntry sends a log entry over WebSocket
func sendLogEntry(logEntry LogEntry, conn *websocket.Conn) {
	// Convert log entry to JSON with "nginx_logs" key
	jsonData, err := json.Marshal(map[string]interface{}{"nginx_logs": logEntry})
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		return
	}

	// Send the JSON data over WebSocket connection
	err = conn.WriteMessage(websocket.TextMessage, jsonData)
	if err != nil {
		log.Println("write:", err)
		return
	}

	// Print the sent log entry
	fmt.Println("Sent log entry:", logEntry)
}
