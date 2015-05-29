package main

import (
	"github.com/gocql/gocql"
)

type Handler struct {
	Messages []*Message
}

func (h *Handler) MessageLogger(server bool, addr *net.UDPAddr, data []byte) {
	m := &Message{server, addr, data}
	h.Messages = append(h.Messages, m)
	fmt.Printf("One: Server:%v, data: %s\n", server, data)
}
