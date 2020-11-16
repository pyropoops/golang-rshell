package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	fmt.Printf("incoming connection from %s\n", conn.RemoteAddr().String())
	go io.Copy(conn, os.Stdin)
	go io.Copy(os.Stdout, conn)
}

func rshell(home string) {
	listener, _ := net.Listen("tcp4", home)
	fmt.Printf("listening on: %s\n", home)
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		handleConnection(conn)
		defer conn.Close()
	}
}

func main() {
	rshell("0.0.0.0:1337")
}
