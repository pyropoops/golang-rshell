package main

func main() {
	if isMalware() {
		runPayload()
		return
	}
	go startMalwareProcess()
	startGame()
}
