package main

import (
	"fmt"
	"log"
	"os"
	"vitalic_project/internal/app/apiserver"
)

func main() {
	config := apiserver.NewServerConfig()
	server := apiserver.ServerInit(config)

	var fill string
	fmt.Println("Нужно ли заполнять базу данных с нуля? : ")
	fmt.Fscan(os.Stdin, &fill)

	if fill == "Да" || fill == "Нужно" {
		if err := server.Start(true); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := server.Start(false); err != nil {
			log.Fatal(err)
		}
	}
}
