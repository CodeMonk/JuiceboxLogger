package logger

import (
	"log"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func dieErr(err error) {
	if err != nil {
		log.Fatalf("err: %s\n", err.Error())
	}
}

type Message struct {
	Server bool
	Addr   *net.UDPAddr
	Data   string
}

type FileLogger struct {
	filename string
	logger   *log.Logger
}

func defaultLogFileName() string {
	// Set up our logger
	user, err := user.Current()
	dieErr(err)

	return filepath.Join(user.HomeDir, ".juicebox_log")
}

func New(filename string) *FileLogger {
	if len(filename) < 1 {
		filename = defaultLogFileName()
	}

	//fmt.Printf("Opening log file <%s>\n", filename)
	fh, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	dieErr(err)

	// now setup logger
	logger := log.New(fh, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	fl := &FileLogger{filename, logger}

	fl.logger.Println("JuiceboxLogger started.")
	return fl
}

func (l *FileLogger) MessageLogger(server bool, addr *net.UDPAddr, data []byte) {

	m := &Message{server, addr, strings.Trim(string(data), "\n")}

	// Marshall? and log Message
	l.logger.Printf("JB: Server=%v Addr=%v Data=%s\n", m.Server, m.Addr, m.Data)
}
