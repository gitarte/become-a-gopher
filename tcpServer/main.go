package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	addr := os.Getenv("ADDR")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	checkErr(err)

	service := addr + ":" + strconv.Itoa(port)
	listener, err := net.Listen("tcp4", service)
	checkErr(err)

	for {
		conn, err := listener.Accept()
		checkErr(err)
		go handler(conn)
	}
}

//	handler - returns cirrent timetamp in required format
func handler(conn net.Conn) {
	//	we ensure that the connection will be closed after this function finished it's execution
	defer conn.Close()

	//	we capture curent timestamp
	now := time.Now()

	//	we read data from client
	buf := make([]byte, 512, 512)
	conn.Read(buf)
	data := string(buf)
	fmt.Printf("Data: %s\n", data)

	//	we translate the data into time format
	result := ""
	switch data {
	case "RFC1123":
		result = now.Format(time.RFC1123)
	case "RFC3339":
		result = now.Format(time.RFC1123)
	default:
		//	custom format
		result = now.Format("2006-01-02 15:04:05.999999999")
	}

	//	we send the tresponse
	conn.Write([]byte(result))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
