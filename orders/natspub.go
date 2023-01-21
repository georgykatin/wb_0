package orders

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"time"
	"wb/Randomaizer"
)

type Pub interface {
	Pub(subject string, data []byte) error
	PubAsync(subject string, data []byte, ah stan.AckHandler) (string, error)
}

type publisher struct {
	stanConn stan.Conn
}

func NewPublisher(stanConn stan.Conn) *publisher {
	return &publisher{stanConn: stanConn}
}

func (p *publisher) Publish(subject string, data []byte) error {
	log.Printf("Publish data: %v to subject: %v", string(data), subject)
	return p.stanConn.Publish(subject, data)
}

func (p *publisher) PublishAsync(subject string, data []byte, ah stan.AckHandler) (string, error) {
	log.Printf("Publish data: %v to subject: %v", string(data), subject)
	return p.stanConn.PublishAsync(subject, data, ah)
}

func (p *publisher) Run() {
	for {

		order := Randomaizer.RandomOrder()
		orderBytes, _ := json.Marshal(*order)

		log.Println("Publish new random order")
		err := p.stanConn.Publish("order:create", orderBytes)

		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(3000 * time.Millisecond)
	}

}
