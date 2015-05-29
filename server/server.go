// Handle emotorwerks client messages.

package server

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"time"
)

type ServerMessage struct {
	Timestamp  time.Time
	ServerTime time.Time
	Id         string // Provided by instantiater
	Data       []*KeyValuePair
}

type KeyValuePair struct {
	Key   string
	Value int
}

func (sm *ServerMessage) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Time        : %v\n", sm.Timestamp))
	buffer.WriteString(fmt.Sprintf("Server Time : %v\n", sm.ServerTime))
	buffer.WriteString(fmt.Sprintf("Id          : %v\n", sm.Id))
	buffer.WriteString(fmt.Sprintf("Data        :\n"))
	for _, kvp := range sm.Data {
		buffer.WriteString(fmt.Sprintf("              %v\n", kvp))
	}
	return buffer.String()
}

func (kvp *KeyValuePair) String() string {
	return fmt.Sprintf("%v : %v", kvp.Key, kvp.Value)
}

// Parse a client message, and return a populated structure
// Sample message: CMD51146A40M60C000$
func New(data string, id string) *ServerMessage {
	sm := &ServerMessage{}
	sm.Timestamp = time.Now()
	sm.Id = id

	err := sm.parseMessage(data)
	if err != nil {
		log.Printf("Error: Unable to parse client string <%s>: %s\n", data, err.Error())
		return nil
	}

	return sm
}

func (sm *ServerMessage) parseMessage(data string) error {

	// parse message
	// Check for CMD in the beginning
	if data[0:3] != "CMD" {
		return fmt.Errorf("Unable to find CMD at start of data: %v\n", data[0:3])
	}
	data = data[3:]

	// Parse out timestamp & Allowed amps
	sm.parseTimeStamp(data[0:5])
	data = data[5:]

	// Finally, our AVPs
	avps := data
	var kvps []*KeyValuePair
	for avps[0] != '$' {
		switch {
		case ('a' <= avps[0] && avps[0] <= 'z') || ('A' <= avps[0] && avps[0] <= 'Z'):
			// Make our avps
			avps = parseAvp(avps, &kvps)
		default:
			return fmt.Errorf("Error:  Illegal token for key value '%c': %s\n", avps[0], avps)
		}
	}
	sm.Data = kvps

	return nil
}

func intToDay(day uint8) string {
	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	return days[int(day)%len(days)]
}

// Parse out timestamp  WHHMMAA == Week, Hour, Minute
func (sm *ServerMessage) parseTimeStamp(data string) {
	day := data[0] - '0'
	hour, err := strconv.Atoi(data[1:3])
	if err != nil {
		log.Printf("Error converting %s to int: %s\n", data[1:3], err.Error())
	}
	minute, err := strconv.Atoi(data[3:5])
	if err != nil {
		log.Printf("Error converting %s to int: %s\n", data[3:5], err.Error())
	}

	timeString := fmt.Sprintf("%s %02d:%02d", intToDay(day), hour, minute)
	sm.ServerTime, err = time.Parse("Mon 15:04", timeString)
	if err != nil {
		log.Printf("Error converting time String: %s: %s\n", timeString, err.Error())
	}
}

func isDigit(char uint8) bool {
	return char >= '0' && char <= '9'
}

func parseAvp(avps string, kvps *[]*KeyValuePair) string {
	key := string(avps[0])
	avps = avps[1:]

	index := 0
	for isDigit(avps[index]) && index < len(avps) {
		index++
	}

	// Wonder what an atoi will do?
	value, err := strconv.Atoi(avps[:index])
	if err != nil {
		log.Fatalf("Unable to parse integer from %s:%s\n", avps[:index], err.Error())
	}
	kvp := &KeyValuePair{key, value}
	*kvps = append(*kvps, kvp)

	for isDigit(avps[0]) {
		avps = avps[1:]
	}
	return avps
}
