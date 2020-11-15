package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	fmt.Println("incoming connection...")
	go io.Copy(conn, os.Stdin)
	go io.Copy(os.Stdout, conn)
}

func rshell(home string) {
	listener, _ := net.Listen("tcp4", home)
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
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("usage: rshell <home_addr>")
		return
	}
	rshell(args[0])
}
