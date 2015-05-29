package main

import (
	"flag"

	"github.com/CodeMonk/JuiceboxLogger/logger"

	"github.com/CodeMonk/UdpProxy/udp"
)

var (
	listenPort = flag.Int("port", 8042, "Port to listen on")
	server     = flag.String("server", "emotorwerks.com", "Destination to proxy to")
	logFile    = flag.String("logfile", "", "File to log traffic to. (~/.juicebox_log default)")
)

func init() {
	flag.Parse()
}

func main() {

	proxy := &udp.UdpProxy{}

	handler := logger.New(*logFile)

	proxy.AddHandler(handler)

	proxy.Run(*listenPort, *server)
}
