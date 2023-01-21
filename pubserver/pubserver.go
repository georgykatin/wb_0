package pubserver

import (
	"context"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"os/signal"
	"syscall"
	"wb/config"
	"wb/orders"
)

const (
	sendOrderSubject = "order:send"
)

type pub_server struct {
	log      *log.Logger
	cfg      *config.Config
	natsConn stan.Conn
}

func NewServer(
	cfg *config.Config,
	natsConn stan.Conn,
) *pub_server {
	return &pub_server{cfg: cfg, natsConn: natsConn}
}

func (ps *pub_server) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		orderPublisher := orders.NewPublisher(ps.natsConn)
		orderPublisher.Run()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Fatalf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Fatalf("ctx.Done: %v", done)
	}

	//log.Println("Server Exited Property")

	return nil

}
