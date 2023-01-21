package subscriber

import (
	"github.com/nats-io/stan.go"
	"log"
	"wb/config"
)

func Subscriber() {
	sub, err := stan.Connect(config.NatsClusterId, "sub")
	log.Printf("Connection to publisher service")
	if err != nil {
		log.Printf("Subscriber can not connect to nats: %v\n", err)
		return
	}
	defer sub.Close()
}
