package models

import (
	"fmt"
	"time"
)

type Order struct {
	OrderUID          string    `json:"order_uid" db:"order_uid"`
	TrackNumber       string    `json:"track_number" db:"track_number"`
	Entry             string    `json:"entry" db:"entry"`
	Delivery          Delivery  `json:"delivery" db:"-"`
	Payment           Payment   `json:"payment" db:"-"`
	Items             []Item    `json:"items" db:"-"`
	Locale            string    `json:"locale" db:"locale"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature"`
	CustomerID        string    `json:"customer_id" db:"customer_id"`
	DeliveryService   string    `json:"delivery_service" db:"delivery_service"`
	ShardKey          string    `json:"shardkey" db:"shardkey"`
	SmID              int       `json:"sm_id" db:"sm_id"`
	DateCreated       time.Time `json:"date_created" db:"date_created"`
	OofShard          string    `json:"oof_shard" db:"oof_shard"`
}

type Delivery struct {
	OrderUID string `json:"-" db:"order_uid"`
	Name     string `json:"name" db:"name"`
	Phone    string `json:"phone" db:"phone"`
	Zip      string `json:"zip" db:"zip"`
	City     string `json:"city" db:"city"`
	Address  string `json:"address" db:"address"`
	Region   string `json:"region" db:"region"`
	Email    string `json:"email" db:"email"`
}

type Payment struct {
	OrderUID     string  `json:"-" db:"order_uid"`
	Transaction  string  `json:"transaction" db:"transaction"`
	RequestID    string  `json:"request_id" db:"request_id"`
	Currency     string  `json:"currency" db:"currency"`
	Provider     string  `json:"provider" db:"provider"`
	Amount       float64 `json:"amount" db:"amount"`
	PaymentDt    int64   `json:"payment_dt" db:"payment_dt"`
	Bank         string  `json:"bank" db:"bank"`
	DeliveryCost float64 `json:"delivery_cost" db:"delivery_cost"`
	GoodsTotal   float64 `json:"goods_total" db:"goods_total"`
	CustomFee    float64 `json:"custom_fee" db:"custom_fee"`
}

type Item struct {
	ID          int64   `json:"-" db:"id"`
	OrderUID    string  `json:"-" db:"order_uid"`
	ChrtID      int64   `json:"chrt_id" db:"chrt_id"`
	TrackNumber string  `json:"track_number" db:"track_number"`
	Price       float64 `json:"price" db:"price"`
	RID         string  `json:"rid" db:"rid"`
	Name        string  `json:"name" db:"name"`
	Sale        float64 `json:"sale" db:"sale"`
	Size        string  `json:"size" db:"size"`
	TotalPrice  float64 `json:"total_price" db:"total_price"`
	NmID        int64   `json:"nm_id" db:"nm_id"`
	Brand       string  `json:"brand" db:"brand"`
	Status      int     `json:"status" db:"status"`
}

type OrderDTO struct {
	OrderUID        string    `json:"order_uid" db:"order_uid"`
	TrackNumber     string    `json:"track_number" db:"track_number"`
	CustomerID      string    `json:"customer_id" db:"customer_id"`
	DeliveryService string    `json:"delivery_service" db:"delivery_service"`
	DateCreated     time.Time `json:"date_created" db:"date_created"`
	TotalPrice      float64   `json:"total_price" db:"total_price"`
	Items           []string  `json:"items" db:"-"`
}

func OrderToDTO(o *Order) OrderDTO {
	fmt.Println(o.Items)
	items := []string{}
	sum := 0
	for _, i := range o.Items {
		items = append(items, i.Name)
		sum += int(i.TotalPrice)
	}
	return OrderDTO{OrderUID: o.OrderUID, TrackNumber: o.TrackNumber,
		CustomerID: o.CustomerID, DeliveryService: o.DeliveryService,
		DateCreated: o.DateCreated, Items: items, TotalPrice: float64(sum)}
}
