package logger

import (
	"bytes"
	"log"
	"net"
	"strings"

	"testing"
)

func TestDefaultLogFilename(t *testing.T) {
	fn := defaultLogFileName()
	if len(fn) < 1 {
		t.Errorf("Length of filename is %d!", len(fn))
	}
	if !strings.Contains(fn, ".juicebox_log") {
		t.Errorf("Filename should contain .juicebox_log, but doesn't <%s>\n", fn)
	}
}

func TestLogger(t *testing.T) {
	// Create byte object
	var b bytes.Buffer // A Buffer needs no initialization.
	logger := log.New(&b, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	fl := &FileLogger{"testing", logger}

	// And send a message
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:3000")
	if err != nil {
		t.Errorf("Error resolving address: %s\n", err.Error())
	}

	fl.MessageLogger(true, addr, []byte("Some Data"))

	if !strings.Contains(b.String(), "Some Data") {
		t.Errorf("Data did not make it to buffer.  Buffer: %s\n", b.String())
	}
}
