// Handle emotorwerks client messages.

package client

import (
	"fmt"
	"testing"
)

// Parse a client message, and return a populated structure
// Sample message: 0812021255561333571420156716:V2468,L93,S0,T25,E40,i87,e0,t30,f5999,m40:\n
var (
	clientMessages = []string{"0812021255561333571420156716:V2467,L93,S0,T25,E40,i83,e20,t30,f5999,m40:\n",
		"0812021255561333571420156716:V2468,L93,S0,T25,E40,i87,e0,t30,f5999,m40:\n"}
)

func TestNewClient(t *testing.T) {
	cm := New(clientMessages[0])
	if cm == nil {
		t.Errorf("Unable to parse client message: %s\n", clientMessages[0])
	}
	fmt.Printf("%v\n", cm)
}
