package model

import (
	"net/http"
)

type Order struct {
	Base
	OrderStatus string `json:"orderStatus"`
	ProductId   string `json:"productId"`
	Num         string `json:"num"`
	BuyerId     string `json:"buyerId"`
}

func OrderGetAll(w http.ResponseWriter, r *http.Request) {

}
func OrderGet(w http.ResponseWriter, r *http.Request) {

}
func OrderAdd(w http.ResponseWriter, r *http.Request) {

}
func OrderPay(w http.ResponseWriter, r *http.Request) {

}
