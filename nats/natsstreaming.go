package nats

import (
	"github.com/nats-io/stan.go"
	"log"
	"time"
	"wb/config"
)

const (
	connectWait        = time.Second * 20
	pubAckWait         = time.Second * 20
	interval           = 5
	maxOut             = 5
	maxPubAcksInflight = 30
)

func NewNatsConnect(cfg *config.Config, clientID string) (stan.Conn, error) {

	return stan.Connect(
		cfg.NatsClusterId,
		clientID,
		stan.ConnectWait(connectWait),
		stan.PubAckWait(pubAckWait),
		stan.NatsURL("nats://"+config.NatsHostName+":4222"),
		stan.Pings(interval, maxOut),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			log.Fatalf("Connection lost, reason: %v", reason)
		}),
		stan.MaxPubAcksInflight(maxPubAcksInflight),
	)
}
