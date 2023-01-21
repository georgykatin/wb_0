package publisher

import (
	"github.com/nats-io/stan.go"
	"log"
	"wb/config"
)

func Publisher() {
	sc, err := stan.Connect(config.NatsClusterId, "publisher")
	log.Printf("Starting publisher service,connect to nats streaming")
	if err != nil {
		log.Printf("Cannot connect to nats streaming: %v\n", err)
		return
	}
	defer sc.Close()

}
