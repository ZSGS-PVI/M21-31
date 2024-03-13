package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/gorilla/websocket"
)

type LogEntry struct {
	Timestamp      string `json:"timestamp"`
	Hostname       string `json:"hostname"`
	Kernel         string `json:"kernel"`
	AuditType      string `json:"audit_type"`
	Operation      string `json:"operation"`
	Profile        string `json:"profile"`
	Unconfined     string `json:"unconfined,omitempty"`
	PID            int    `json:"pid"`
	Comm           string `json:"comm"`
	Capability     string `json:"capability,omitempty"`
	Capname        string `json:"capname,omitempty"`
	Info           string `json:"info,omitempty"`
	AuditEpoch     string `json:"audit_epoch"`
	TimestampEpoch string `json:"timestamp_epoch"`
}

type WebSocketData struct {
	KvmLogs []LogEntry `json:"kvm_logs"`
}

func main() {
	// WebSocket connection to Java servlet
	log.Println("Connecting to WebSocket server...")
	conn, _, err := websocket.DefaultDialer.Dial("ws://vishnu:8087/RedisWebSoc/serverws", nil)
	if err != nil {
		log.Fatal("Failed to connect to WebSocket server:", err)
	}
	defer conn.Close()
	log.Println("Successfully connected to WebSocket server.")

	// Open the log file
	log.Println("Opening log file...")
	file, err := os.Open("/var/log/kern.log.1")
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer file.Close()
	log.Println("Log file opened successfully.")

	// Read log entries from the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		entry := parseLogEntry(line)
		if entry != nil {
			data := map[string]interface{}{
				"kvm_logs": entry,
			}

			// Convert LogEntry to JSON format and print it to console
			jsonData, err := json.Marshal(data)
			if err != nil {
				log.Fatal("Failed to marshal JSON data:", err)
			}
			fmt.Println("Log Entry in JSON Format:", string(jsonData))

			// Send the JSON data over WebSocket connection
			log.Println("Sending data to WebSocket server...")
			err = conn.WriteMessage(websocket.TextMessage, jsonData)
			if err != nil {
				log.Println("Failed to send data to WebSocket server:", err)
			} else {
				log.Println("Data sent successfully to WebSocket server.")
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading log file:", err)
	}
}

func parseLogEntry(line string) *LogEntry {
	re := regexp.MustCompile(`^(?P<timestamp>\w+\s+\d+\s\d+:\d+:\d+)\s(?P<hostname>\S+)\s(?P<kernel>\S+):\s\[\s*(?P<timestamp_epoch>\d+\.\d+)\]\saudit:\s+type=(?P<audit_type>\d+)\s+audit\((?P<audit_epoch>\d+\.\d+):\d+\):\s+apparmor="(?P<operation>\w+)"\s+operation="(?P<profile>\w+)"(?:\s+info="(?P<info>[^"]+)")?\s+profile="(?P<unconfined>\w+)"\s+pid=(?P<pid>\d+)\s+comm="(?P<comm>[^"]+)"(?:\s+capability=(?P<capability>\d+)\s+capname="(?P<capname>\w+)")?`)

	match := re.FindStringSubmatch(line)
	if match == nil {
		return nil
	}

	groups := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			groups[name] = match[i]
		}
	}

	pid, err := strconv.Atoi(groups["pid"])
	if err != nil {
		log.Println("Failed to parse PID:", err)
		return nil
	}

	return &LogEntry{
		Timestamp:      groups["timestamp"],
		Hostname:       groups["hostname"],
		Kernel:         groups["kernel"],
		AuditType:      groups["audit_type"],
		Operation:      groups["operation"],
		Profile:        groups["profile"],
		Unconfined:     groups["unconfined"],
		PID:            pid,
		Comm:           groups["comm"],
		Capability:     groups["capability"],
		Capname:        groups["capname"],
		Info:           groups["info"],
		AuditEpoch:     groups["audit_epoch"],
		TimestampEpoch: groups["timestamp_epoch"],
	}
}
