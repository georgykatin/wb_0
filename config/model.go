package config

import "time"

type Order struct {
	OrderUid          string    `json:"order_uid" validate:"required" fake:"{regex:[0-9a-z]{19}}"`
	TrackNumber       string    `json:"track_number" validate:"required" fake:"{regex:[A-Z]{14}}"`
	Entry             string    `json:"entry" validate:"required" fake:"{regex:[A-Z]{4}}"`
	Locale            string    `json:"locale" validate:"required" fake:"{randomstring:[ru,en,it]}"`
	InternalSignature string    `json:"internal_signature" validate:"required" fake:"{lettern:8}"`
	CustomerId        string    `json:"customer_id" validate:"required"  fake:"{lettern:4}"`
	DeliveryService   string    `json:"delivery_service" validate:"required" fake:"{lettern:5}"`
	Shardkey          string    `json:"shardkey" validate:"required" fake:"{regex:[0-9]{2}}"`
	SmId              int       `json:"sm_id" validate:"required" fake:"{number:1, 5000}"`
	DateCreated       time.Time `json:"date_created" validate:"required" fake:"{year}-{month}-{day}T{hour}:{minute}:{second}Z" format:"2006-01-02T06:22:19Z"`
	OofShard          string    `json:"oof_shard" validate:"required" fake:"{regex:[0-9]{2}}"`
	Delivery          `json:"delivery"`
	Payment           `json:"payment"`
	Items             []Items `json:"items"`
}

type Delivery struct {
	Name    string `json:"name" validate:"required" fake:"{firstname} {lastname}"`
	Phone   string `json:"phone" validate:"required" fake:"{phone}"`
	Zip     string `json:"zip" validate:"required" fake:"{zip}"`
	City    string `json:"city" validate:"required" fake:"{city}"`
	Address string `json:"address" validate:"required" fake:"{streetname}"`
	Region  string `json:"region" validate:"required"  fake:"{state}"`
	Email   string `json:"email" validate:"required" fake:"{email}"`
}

type Payment struct {
	Transaction  string `json:"transaction" validate:"required" fake:""`
	RequestId    string `json:"requestId" validate:"required" fake:""`
	Currency     string `json:"currency" validate:"required" fake:""`
	Provider     string `json:"provider" validate:"required" fake:""`
	Amount       int    `json:"amount" validate:"required" fake:"{number:1, 2389212}"`
	PaymentDt    int    `json:"paymentDt" validate:"required" fake:"{number:1, 2389212}"`
	Bank         string `json:"bank" validate:"required" fake:""`
	DeliveryCost int    `json:"deliveryCost" validate:"required" fake:"{number:1, 2389212}"`
	GoodTotal    int    `json:"goodTotal" validate:"required" fake:"{number:1, 2389212}"`
	CustomFee    int    `json:"customFee" validate:"required" fake:"{number:1, 2389212}"`
}

type Items struct {
	ChrtId      int    `json:"chrtId" validate:"required" fake:"{number:1, 5000}"`
	TrackNumber string `json:"trackNumber" validate:"required" fake:"{regex:[A-Z]{14}}"`
	Price       int    `json:"price" validate:"required" fake:"{number:1, 5000000}"`
	Rid         string `json:"rid" validate:"required" fake:"{regex:[0-9a-z]{21}}"`
	Name        string `json:"name" validate:"required" fake:"{minecraftarmorpart}"`
	Sale        int    `json:"sale" validate:"required" fake:"{number:1, 100}"`
	Size        string `json:"size" validate:"required" fake:"{regex:[0-1]{3}}"`
	TotalPrice  int    `json:"totalPrice" validate:"required" fake:"{number:1, 5000000}"`
	NmId        int    `json:"nmId" validate:"required" fake:"{number:1, 2389212}"`
	Brand       string `json:"brand" validate:"required" fake:"{firstname} {lastname}"`
	Status      int    `json:"status" validate:"required" fake:"{regex:[0-1]{3}}"`
}
