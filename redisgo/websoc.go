package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

type RedisLogEntry struct {
	LogEntry LogEntry `json:"redis_log"`
}

type LogEntry struct {
	Timestamp string     `json:"timestamp"`
	IPDetails AddrDetail `json:"addr_details"`
	Command   string     `json:"command"`
	Content   Content    `json:"content"`
}
type AddrDetail struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

type Content struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Salary int    `json:"salary"`
	Dept   string `json:"dept"`
}

func main() {
	cmd := exec.Command("redis-cli", "monitor")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("Error creating pipe:", err)
		return
	}

	err = cmd.Start()
	if err != nil {
		log.Println("Error starting command:", err)
		return
	}

	// WebSocket connection to Java servlet
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8087/RedisWebSoc/serverws", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	count := 0

	scanner := bufio.NewScanner(stdout)
	//	var addrSlice []AddrDetail
	for scanner.Scan() {
		line := scanner.Text()
		if line == "OK" {
			continue
		}

		parts := strings.Split(line, " ")
		if len(parts) < 4 {
			log.Println("Invalid line:", line)
			continue
		}

		timestamp := parts[0]
		ipDetails := strings.Trim(parts[2], "[]")
		addr := strings.Split(ipDetails, ":")
		port, _ := strconv.Atoi(addr[1])
		//addrSlice = append(addrSlice, AddrDetail{addr[0], port})

		command := parts[3]
		jsonStr := parts[len(parts)-1]

		jsonStr = strings.Trim(jsonStr, "\"")
		jsonStr = strings.ReplaceAll(jsonStr, "\\\"", "\"")

		var content Content
		err := json.Unmarshal([]byte(jsonStr), &content)
		if err != nil {
			log.Printf("Error unmarshalling JSON: %v, JSON string: %s\n", err, jsonStr)
			continue
		}

		// Convert timestamp to a float64
		timestampFloat, err := strconv.ParseFloat(timestamp, 64)
		if err != nil {
			log.Printf("Error converting timestamp to float64: %v, Timestamp: %s\n", err, timestamp)
			continue
		}

		// Convert Unix timestamp to a time.Time value
		t := time.Unix(int64(timestampFloat), int64((timestampFloat-float64(int64(timestampFloat)))*1e9))

		// Format the time with fractional seconds
		formattedTime := t.Format("2006-01-02 15:04:05.999999999")

		// Print the formatted time
		//fmt.Println("Formatted time:", formattedTime)
		logEntry := LogEntry{
			Timestamp: formattedTime,
			IPDetails: AddrDetail{addr[0], port},
			Command:   command,
			Content:   content,
		}

		// Create a RedisLogEntry instance
		redisLogEntry := RedisLogEntry{
			LogEntry: logEntry,
		}

		// Marshal the RedisLogEntry into JSON
		jsonData, err := json.Marshal(redisLogEntry)
		if err != nil {
			log.Fatalf("Error marshalling JSON: %v", err)
		}

		fmt.Println(string(jsonData))
		//fmt.Println(count)
		count++
		// Send the JSON data over WebSocket connection
		err = conn.WriteMessage(websocket.TextMessage, jsonData)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scanner error: %v", err)
	}
}
