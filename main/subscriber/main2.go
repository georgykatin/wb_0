package main

import (
	"log"
	"wb/cache"
	config "wb/config"
	"wb/httpserver"
	"wb/nats"
	"wb/orders"
)

func main() {

	var (
		config2 config.Config
		ch      cache.Cache
	)
	natsConn, err := nats.NewNatsConnect(&config2, "subscriber")
	if err != nil {
		log.Fatalf("Error fail to init nats sub connection")
	}

	pgxPool, err := orders.NewPgxConn(&config2)
	if err != nil {
		log.Fatalf("Error fail to connect db main file %v", err)
	}
	serv2 := httpserver.NewServer(&config2, natsConn, pgxPool, &ch)

	log.Fatal(serv2.Run())

}
