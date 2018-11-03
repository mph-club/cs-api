package main

import (
	"csportal-server/client"
	"csportal-server/server"
)

func main() {
	client.CreateUserThroughAWS()

	server.CreateAndListen()
}
