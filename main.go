package main

import (
	"flag"

	"github.com/CodeMonk/JuiceboxLogger/logger"

	"github.com/CodeMonk/UdpProxy/udp"
)

var (
	listenPort = flag.Int("port", 8042, "Port to listen on")
	server     = flag.String("server", "emotorwerks.com", "Destination to proxy to")
	logFile    = flag.String("logfile", "Juicebox.log", "File to log traffic to")
)

func init() {
	flag.Parse()
}

func main() {

	proxy := &udp.UdpProxy{}

	handler := &logger.FileLogger{}

	proxy.AddHandler(handler)

	proxy.Run(*listenPort, *server)
}
