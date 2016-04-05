package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func requestHandler(conn net.Conn) {
	println("Message received")

	msg, endRequest := processMsg(conn)

	if endRequest {
		conn.Write([]byte(msg))
		conn.Close()
	}
}

func processMsg(conn net.Conn) (msg string, endRequest bool) {
	reply := make([]byte, 1024)

	_, err := conn.Read(reply)

	if err != nil {
		os.Exit(1)
	}

	replyStr := string(reply)

	endRequest = strings.Contains(replyStr, "END")
	msg = strings.Replace(replyStr, "\nEND\n", "", -1)

	return
}

func main() {
	listen, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		os.Exit(1)
	}

	defer listen.Close()
	fmt.Println("Listening on localhost:1234")

	for {
		conn, err := listen.Accept()
		if err != nil {
			os.Exit(1)
		}

		go requestHandler(conn)
	}
}
