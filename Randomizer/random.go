package Randomaizer

import (
	"github.com/brianvoe/gofakeit"
	"log"
	"wb/config"
)

func RandomOrder() *config.Order {
	var Order config.Order
	gofakeit.Seed(0)
	gofakeit.Struct(Order)
	log.Println(Order)
	return &Order
}
