package dilivery

import (
	"context"
	"encoding/json"
	"log"
	"sync"

	"wb/Using"
	"wb/config"

	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
)

const (
	createOrderWorkers = 0
	createOrderSubject = "order:create"
	orderGroupName     = "order_service"
)

type OrderSub struct {
	stanConn     stan.Conn
	orderUseCase Using.UseCase
	validate     *validator.Validate
}

func (sub *OrderSub) NewOrderSub(stanConn stan.Conn, orderUseCase Using.UseCase, validate *validator.Validate) *OrderSub {
	return &OrderSub{stanConn: stanConn, orderUseCase: orderUseCase, validate: validate}
}

func (sub *OrderSub) Subscribe(subject, qgroup string, workerNum int, cb stan.MsgHandler) {
	log.Printf("Subscribing to Subject: %v, group: %v", subject, qgroup)
	wg := &sync.WaitGroup{}

	for i := 0; i <= workerNum; i++ {
		wg.Add(1)
		go sub.runWorker(
			wg,
			i,
			sub.stanConn,
			subject,
			qgroup,
			cb,
		)
	}
	wg.Wait()
}

func (sub *OrderSub) runWorker(
	wg *sync.WaitGroup,
	workerID int,
	conn stan.Conn,
	subject string,
	qgroup string,
	cb stan.MsgHandler,
	opts ...stan.SubscriptionOption,
) {
	log.Printf("Subscribing worker: %v, subject: %v, qgroup: %v", workerID, subject, qgroup)
	defer wg.Done()

	_, err := conn.QueueSubscribe(subject, qgroup, cb, opts...)
	if err != nil {
		log.Printf("WorkerID: %v, QueueSubscribe: %v", workerID, err)
		if err := conn.Close(); err != nil {
			log.Printf("WorkerID: %v, conn.Close error: %v", workerID, err)
		}
	}

}

func (sub *OrderSub) Run(ctx context.Context) {

	go sub.Subscribe(createOrderSubject, orderGroupName, createOrderWorkers, sub.processCreateOrder(ctx))
}

func (sub *OrderSub) processCreateOrder(ctx context.Context) stan.MsgHandler {
	return func(msg *stan.Msg) {
		log.Printf("subscriber process Create Order: %s", msg.String())

		var m config.Order

		err := sub.validate.Struct(m)

		if err != nil {
			log.Print("Data validate error")
			if _, ok := err.(*validator.InvalidValidationError); ok {
				log.Println(err)
				return
			}
		}

		if err := json.Unmarshal(msg.Data, &m); err != nil {
			log.Printf("json.Unmarshal: %v", err)
			return
		}

		// if err := s.orderUC.Create(ctx, &m); err != nil {
		// 	s.log.Printf("orderUC.Create : %v", err)
		// 	return
		// }

		if err := sub.orderUseCase.BatchCreate(ctx, m); err != nil {
			log.Printf("orderUC.Create : %v", err)
			return
		}
	}
}
