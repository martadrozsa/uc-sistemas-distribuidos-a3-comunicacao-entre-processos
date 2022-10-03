package main

import (
	"encoding/json"
	_ "encoding/json"
	"github.com/martadrozsa/uc-sistemas-distribuidos-a3-comunicacao-entre-processos/dto"
	"log"
	"net"
)

const ADDRESS = "localhost:8080"

func main() {

	clientConnection, err := net.Dial("tcp", ADDRESS)
	checkError(err)

	requestOne := dto.Request{Ids: []int{10, 20, 30, 40, 50, 60, 70}}

	encoder := json.NewEncoder(clientConnection)
	err = encoder.Encode(requestOne)

	decoder := json.NewDecoder(clientConnection)

	var res dto.Response
	err = decoder.Decode(&res)
	checkError(err)

	log.Printf("Total R$%.2f", res.TotalPrice)

	_ = clientConnection.Close()

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
