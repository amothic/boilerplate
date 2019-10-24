package main

import "log"

func main() {
	server, err := InitializeServer()
	if err != nil {
		log.Fatal(err)
	}
	server.Run()
}
