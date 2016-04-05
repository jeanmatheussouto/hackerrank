package main

import (
	"fmt"
	"net"
	"testing"
)

var msgsTest = []struct {
	in  string
	out string
}{
	{"Client 1:\nHello World\nEND", "Client 1:\nHello World\n"},
	{"Client 2:\nThis is line 1\nThis is line 2\nEND", "Client 2:\nThis is line 1\nThis is line 2"},
}

func TestConn(t *testing.T) {
	message := "Hi there!"

	go func() {
		conn, err := net.Dial("tcp", ":3000")
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		if _, err := fmt.Fprintf(conn, message); err != nil {
			t.Fatal(err)
		}
	}()

	l, err := net.Listen("tcp", ":3000")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		defer conn.Close()

		conn.Write([]byte(message))

		msg, _ := processMsg(conn)
		print(msg)
		print(message)
		if msg != message {
			t.Fatalf("\nExpected:\t%sGot:\t\t%s", message, msg)
		}
		return // Done
	}

}
