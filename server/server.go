package main

import (
	"bufio"
	"log"
	"net"
)

const SERVER_ADDRESS = "localhost:8080"

func main() {
	listener, err := net.Listen("tcp", SERVER_ADDRESS)
	checkError(err)

	log.Println("Server listening on", SERVER_ADDRESS)

	for {
		serverConnection, err := listener.Accept()
		checkError(err)

		log.Println("Server received a new connection")

		//goroutine
		go handlerConnection(serverConnection)

	}

}

func handlerConnection(connection net.Conn) {
	response, err := bufio.NewReader(connection).ReadString('\n')
	checkError(err)

	//Processamento
	log.Println("Word received:", response)

	_, err = connection.Write([]byte(response))
	checkError(err)

	_ = connection.Close()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
