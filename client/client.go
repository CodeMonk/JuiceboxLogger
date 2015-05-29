// Handle emotorwerks client messages.

package client

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type ClientMessage struct {
	Timestamp time.Time
	Id        string
	Data      []*KeyValuePair
}

type KeyValuePair struct {
	Key   string
	Value int
}

func (cm *ClientMessage) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Time    : %v\n", cm.Timestamp))
	buffer.WriteString(fmt.Sprintf("Id      : %v\n", cm.Id))
	buffer.WriteString(fmt.Sprintf("Data    :\n"))
	for _, kvp := range cm.Data {
		buffer.WriteString(fmt.Sprintf("          %v\n", kvp))
	}
	return buffer.String()
}

func (kvp *KeyValuePair) String() string {
	return fmt.Sprintf("%v : %v", kvp.Key, kvp.Value)
}

// Parse a client message, and return a populated structure
// Sample message: 0812021255561333571420156716:V2468,L93,S0,T25,E40,i87,e0,t30,f5999,m40:\n
func New(data string) *ClientMessage {
	cm := &ClientMessage{}
	cm.Timestamp = time.Now()

	err := cm.parseMessage(data)
	if err != nil {
		log.Printf("Error: Unable to parse client string <%s>: %s\n", data, err.Error())
		return nil
	}

	return cm
}

func (cm *ClientMessage) parseMessage(data string) error {

	// parse message
	colonStrings := strings.Split(data, ":")
	if len(colonStrings) != 3 {
		return fmt.Errorf("Unable to split data into colon payloads: %s\n", data)
	}
	cm.Id = colonStrings[0]

	commaStrings := strings.Split(colonStrings[1], ",")
	kvps := make([]*KeyValuePair, len(commaStrings))
	for i, str := range commaStrings {
		value, err := strconv.Atoi(str[1:])
		if err != nil {
			log.Printf("Error: Unable to find integer value in <%s>: %s\n", str[1:], err.Error())
			value = -1
		}
		kvp := &KeyValuePair{
			Key:   string(str[0]),
			Value: value}
		kvps[i] = kvp
	}
	cm.Data = kvps

	return nil
}
