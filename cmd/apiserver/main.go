package main

import (
	"log"
	"vitalic_project/internal/app/apiserver"
)

func main() {
	config := apiserver.NewServerConfig()
	server := apiserver.ServerInit(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
