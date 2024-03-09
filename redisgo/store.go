// package main

// import (
// 	"bufio"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os/exec"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/gorilla/websocket"
// )

// type RedisLogEntry struct {
// 	LogEntry LogEntry `json:"redis_log"`
// }

// type LogEntry struct {
// 	Timestamp string     `json:"timestamp"`
// 	IPDetails AddrDetail `json:"addr_details"`
// 	Command   string     `json:"command"`
// 	Content   Content    `json:"content"`
// }
// type AddrDetail struct {
// 	IP   string `json:"ip"`
// 	Port int    `json:"port"`
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
// 		ReadBufferSize:   12000000,
// 		WriteBufferSize:  12000000,
// 		Proxy:            nil,
// 		TLSClientConfig:  nil,
// 	}

// 	// WebSocket connection to Java servlet
// 	conn, _, err := dialer.Dial("ws://localhost:8087/RedisWebSoc/serverws", nil)
// 	if err != nil {
// 		log.Fatal("dial:", err)
// 	}
// 	defer conn.Close()

// 	scanner := bufio.NewScanner(stdout)
// 	//	var addrSlice []AddrDetail
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		if line == "OK" {
// 			continue
// 		}

// 		parts := strings.Split(line, " ")
// 		if len(parts) < 4 {
// 			log.Println("Invalid line:", line)
// 			continue
// 		}

// 		timestamp := parts[0]
// 		ipDetails := strings.Trim(parts[2], "[]")
// 		addr := strings.Split(ipDetails, ":")
// 		port, _ := strconv.Atoi(addr[1])
// 		//addrSlice = append(addrSlice, AddrDetail{addr[0], port})

// 		command := parts[3]
// 		jsonStr := parts[len(parts)-1]

// 		jsonStr = strings.Trim(jsonStr, "\"")
// 		jsonStr = strings.ReplaceAll(jsonStr, "\\\"", "\"")

// 		var content Content
// 		err := json.Unmarshal([]byte(jsonStr), &content)
// 		if err != nil {
// 			log.Printf("Error unmarshalling JSON: %v, JSON string: %s\n", err, jsonStr)
// 			continue
// 		}

// 		timestampParts := strings.Split(timestamp, ".")
// 		if len(timestampParts) != 2 {
// 			log.Printf("Invalid Unix timestamp: %s\n", timestamp)
// 			continue
// 		}

// 		secondsStr := timestampParts[0]
// 		millisecondsStr := timestampParts[1]

// 		// Parse seconds part
// 		seconds, err := strconv.ParseInt(secondsStr, 10, 64)
// 		if err != nil {
// 			log.Printf("Error parsing seconds: %v\n", err)
// 			continue
// 		}

// 		// Parse milliseconds part
// 		milliseconds, err := strconv.ParseInt(millisecondsStr, 10, 64)
// 		if err != nil {
// 			log.Printf("Error parsing milliseconds: %v\n", err)
// 			continue
// 		}

// 		// Combine seconds and milliseconds to create Unix timestamp in nanoseconds
// 		unixTimestamp := seconds*1000000000 + milliseconds*1000000

// 		// Convert Unix timestamp to time.Time
// 		timestampTime := time.Unix(0, unixTimestamp)

// 		// Print timestamp in desired format
// 		fmt.Println("Timestamp:", timestampTime.Format("2006/01/02 15:04:05.999999999"))
// 		logEntry := LogEntry{
// 			Timestamp: timestampTime.Format("2006/01/02 15:04:05.999999999"),
// 			IPDetails: AddrDetail{addr[0], port},
// 			Command:   command,
// 			Content:   content,
// 		}

// 		// Create a RedisLogEntry instance
// 		redisLogEntry := RedisLogEntry{
// 			LogEntry: logEntry,
// 		}

// 		// Marshal the RedisLogEntry into JSON
// 		jsonData, err := json.Marshal(redisLogEntry)
// 		if err != nil {
// 			log.Fatalf("Error marshalling JSON: %v", err)
// 		}

// 		fmt.Println(string(jsonData))

// 		fmt.Println(1)

// 		// Send the JSON data over WebSocket connection
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
