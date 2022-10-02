package main

import (
	"io"
	"log"
	"net"
)

const ADDRESS = "localhost:8080"

func main() {

	clientConnection, err := net.Dial("tcp", ADDRESS)
	checkError(err)

	_, err = clientConnection.Write([]byte("teste\n"))
	checkError(err)

	response, err := io.ReadAll(clientConnection)
	checkError(err)

	log.Println(string(response))

	_ = clientConnection.Close()

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
