// package main

// import (
// 	"bufio"
// 	"encoding/json"
// 	"log"
// 	"os/exec"
// 	"strings"
// 	"time"

// 	"github.com/gorilla/websocket"
// )

// type LogEntry struct {
// 	Timestamp string  `json:"timestamp"`
// 	IPDetails string  `json:"ip_details"`
// 	Command   string  `json:"command"`
// 	Content   Content `json:"content"`
// }

// type Content struct {
// 	ID     string `json:"id"`
// 	Name   string `json:"name"`
// 	Salary int    `json:"salary"`
// 	Dept   string `json:"dept"`
// }

// func main() {
// 	cmd := exec.Command("redis-cli", "monitor")

// 	stdout, err := cmd.StdoutPipe()
// 	if err != nil {
// 		log.Println("Error creating pipe:", err)
// 		return
// 	}

// 	err = cmd.Start()
// 	if err != nil {
// 		log.Println("Error starting command:", err)
// 		return
// 	}

// 	dialer := websocket.Dialer{
// 		HandshakeTimeout: 10 * time.Second,
// 		ReadBufferSize:   8192,
// 		WriteBufferSize:  8192,
// 		Proxy:            nil,
// 		TLSClientConfig:  nil,
// 	}

// 	// WebSocket connection to Java servlet
// 	conn, _, err := dialer.Dial("ws://localhost:8087/RedisWebSoc/writerws", nil)
// 	if err != nil {
// 		log.Fatal("dial:", err)
// 	}
// 	defer conn.Close()

// 	scanner := bufio.NewScanner(stdout)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		log.Println(line)
// 		if line == "OK" {
// 			continue
// 		}

// 		parts := strings.Split(line, " ")
// 		if len(parts) < 4 {
// 			log.Println("Invalid line:", line)
// 			continue
// 		}

// 		timestamp := parts[0]
// 		ipDetails := strings.Join(parts[1:3], " ")
// 		command := parts[3]
// 		jsonStr := parts[len(parts)-1]

// 		jsonStr = strings.Trim(jsonStr, "\"")
// 		jsonStr = strings.ReplaceAll(jsonStr, "\\\"", "\"")

// 		var content Content
// 		err := json.Unmarshal([]byte(jsonStr), &content)
// 		if err != nil {
// 			log.Printf("Error unmarshalling JSON: %v, JSON string: %s\n", err, jsonStr)
// 			continue // Skip to the next line
// 		}

// 		logEntry := LogEntry{
// 			Timestamp: timestamp,
// 			IPDetails: ipDetails,
// 			Command:   command,
// 			Content:   content,
// 		}

// 		// Convert LogEntry struct to JSON
// 		jsonData, err := json.Marshal(logEntry)
// 		if err != nil {
// 			log.Fatalf("Error marshalling JSON: %v", err)
// 		}

// 		// Send JSON data over WebSocket connection
// 		err = conn.WriteMessage(websocket.TextMessage, jsonData)
// 		if err != nil {
// 			log.Println("write:", err)
// 			break
// 		}
// 	}
// 	if err := scanner.Err(); err != nil {
// 		log.Fatalf("Scanner error: %v", err)
// 	}
// }
