package models

import "time"

type Order struct {
	OrderID     uint      `json:"order_id"`
	UserID      uint      `json:"user_id"`
	ProductID   uint      `json:"product_id"`
	OrderTotal  float64   `json:"order_total"`
	OrderDate   time.Time `json:"order_date"`
	OrderStatus string    `json:"order_status"`
}
