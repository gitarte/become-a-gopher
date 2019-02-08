package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr := os.Getenv("ADDR")
	port := os.Getenv("PORT")

	value := ""
	if len(os.Args) >= 2 {
		value = os.Args[1]
	}

	service := addr + ":" + port
	client, err := net.Dial("tcp4", service)
	checkErr(err)
	client.Write([]byte(value))
	fmt.Println(len(os.Args))
	buf := make([]byte, 512, 512)
	client.Read(buf)
	fmt.Println(string(buf))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
