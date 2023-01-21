package main

import (
	"log"
	config2 "wb/config"
	"wb/nats"
	"wb/pubserver"
)

func main() {
	var config config2.Config
	natsConnection, err := nats.NewNatsConnect(&config, "publisher")
	if err != nil {
		log.Fatalf("Error fail to init nats pub connection")
	}

	serv1 := pubserver.NewServer(&config, natsConnection)
	log.Fatal(serv1.Run())
}
