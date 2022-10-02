package main

import (
	"encoding/json"
	"github.com/martadrozsa/uc-sistemas-distribuidos-a3-comunicacao-entre-processos/dto"
	"log"
	"net"
)

const ServerAddress = "localhost:8080"

var productsMap = map[int]float64{10: 55.00, 20: 25.00, 30: 35.00, 40: 45.00}

func main() {
	listener, err := net.Listen("tcp", ServerAddress)
	checkError(err)

	log.Println("Server listening on", ServerAddress)

	for {
		serverConnection, err := listener.Accept()
		checkError(err)

		log.Println("Server received a new connection")

		//goroutine
		go handleConnection(serverConnection)
	}
}

func handleConnection(connection net.Conn) {

	decoder := json.NewDecoder(connection)
	encoder := json.NewEncoder(connection)

	var req dto.Request
	err := decoder.Decode(&req)
	checkError(err)

	sumProducts := 0.0
	for _, id := range req.Ids {
		sumProducts += productsMap[id]
	}
	//Processamento
	log.Println("Ids received:", req)

	response := dto.Response{TotalPrice: sumProducts}
	err = encoder.Encode(response)

	checkError(err)

	_ = connection.Close()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
