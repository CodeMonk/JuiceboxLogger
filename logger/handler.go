package logger

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/user"
	"path/filepath"
)

func dieErr(err error) {
	if err != nil {
		log.Fatalf("err: %s\n", err.Error())
	}
}

func init() {
	// Set up our logger
	user, err := user.Current()
	dieErr(err)

	path := filepath.Join(user.HomeDir, ".juicebox_log")
	fmt.Printf("Opening log file <%s>\n", path)
	fh, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	dieErr(err)

	// now setup logger
	log.SetOutput(fh)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	log.Printf("JuiceboxLogger started.\n")
}

type FileLogger struct {
}

type Message struct {
	Server bool
	Addr   *net.UDPAddr
	Data   string
}

func (l *FileLogger) MessageLogger(server bool, addr *net.UDPAddr, data []byte) {

	m := &Message{server, addr, string(data)}

	// Marshall? and log Message
	log.Printf("JB: Server=%v Addr=%v Data=%s\n", m.Server, m.Addr, m.Data)
}
