// Handle emotorwerks server messages.

package server

import (
	"fmt"
	"testing"
)

// Parse a server message, and return a populated structure
// Sample message: CMD51146A40M60C000$
var (
	serverMessages = []string{"CMD51146A40M60C000$", "CMD51147A40M60C000$"}
)

func TestNewServer(t *testing.T) {
	sm := New(serverMessages[0], "1234567890123456")
	if sm == nil {
		t.Errorf("Unable to parse server message: %s\n", serverMessages[0])
	}
	fmt.Printf("%v\n", sm)
}
