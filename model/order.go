package model

import (
	"github.com/StevenZack/tools"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
)

type Order struct {
	Base
	OrderId     string `json:"orderId"`
	OrderStatus string `json:"orderStatus"`
	CreateTime  string `json:"createTime"`
	ProductId   string `json:"productId"`
	Num         int    `json:"num"`
	Price       int    `json:"price"`
	BuyerId     string `json:"buyerId"`
	ProductName string `json:"productName"`
}

func OrderGetAll(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	openId := r.FormValue("openId")
	orderStatus := r.FormValue("orderStatus")
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	co := s.DB(db).C(corder)
	var backData struct {
		Base
		Orders []Order `json:"orders"`
	}
	if orderStatus == "" {
		e = co.Find(bson.M{"buyerid": openId}).All(&backData.Orders)
	} else {
		e = co.Find(bson.M{"buyerid": openId, "orderStatus": orderStatus}).All(&backData.Orders)
	}
	if e != nil {
		returnErr(w, e)
		return
	}
	backData.Status = "OK"
	returnData(w, backData)
}
func OrderGet(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	orderId := r.FormValue("orderId")
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	co := s.DB(db).C(corder)
	order := Order{}
	e = co.Find(bson.M{"orderid": orderId}).One(&order)
	if e != nil {
		returnErr(w, e)
		return
	}
	order.Status = "OK"
	returnData(w, order)
}
func OrderAdd(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	order := Order{}
	order.ProductId = r.FormValue("productId")
	var e error
	order.Num, e = strconv.Atoi(r.FormValue("num"))
	if e != nil {
		returnErr(w, e)
		return
	}
	order.BuyerId = r.FormValue("buyerId")
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	co := s.DB(db).C(corder)
	cu := s.DB(db).C(cuser)
	cp := s.DB(db).C(cproduct)
	count, e := cu.Find(bson.M{"openid": order.BuyerId}).Count()
	if e != nil {
		returnErr(w, e)
		return
	}
	if count < 1 {
		returnErr(w, "账户不存在，请重新登录")
		return
	}
	produc := Product{}
	e = cp.Find(bson.M{"productid": order.ProductId}).One(&produc)
	if e != nil {
		returnErr(w, e)
		return
	}
	if count < 1 {
		returnErr(w, "商品不存在")
		return
	}
	order.OrderStatus = "未付款"
	order.OrderId = tools.NewToken()
	order.CreateTime = tools.GetTimeStrNow()
	order.ProductName = produc.Name
	order.Price = order.Num * produc.Price
	e = co.Insert(order)
	if e != nil {
		returnErr(w, e)
		return
	}
	returnData(w, Base{Status: "OK"})
}
func OrderPay(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	orderId := r.FormValue("orderId")
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	co := s.DB(db).C(corder)
	order := Order{}
	e = co.Find(bson.M{"orderid": orderId}).One(&order)
	if e != nil {
		returnErr(w, e)
		return
	}
	order.OrderStatus = "待收货"
	e = co.Update(bson.M{"orderid": orderId}, order)
	if e != nil {
		returnErr(w, e)
		return
	}
	returnData(w, Base{Status: "OK"})
}
