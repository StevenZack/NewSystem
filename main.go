package main

import (
	"fmt"
	"github.com/StevenZack/NewSystem/model"
	"net/http"
)

func main() {
	http.HandleFunc("/UserServlet/login", model.UserLogin)

	http.HandleFunc("/MomentServlet/upload", model.MomentUpload)
	http.HandleFunc("/MomentServlet/findAll", model.MomentQuery)

	http.HandleFunc("/NewsServlet/findAllByKind", model.NewsFindByType)
	http.HandleFunc("/NewsServlet/findAll", model.NewsFindAll)
	http.HandleFunc("/NewsServlet/findPhotoBrowserByNewsId", model.NewsGetDetail)
	http.HandleFunc("/NewsServlet/add", model.NewsAdd)

	http.HandleFunc("/SearchServlet/findproductByname", model.ProFindByName)
	http.HandleFunc("/SearchServlet/findproductByname_English", model.ProFindByEN)
	http.HandleFunc("/SearchServlet/findproductBytype", model.ProFindByType)
	http.HandleFunc("/SearchServlet/fuzzy", model.ProFuzzyFind)

	http.HandleFunc("/ProServlet/add", model.ProAdd)

	http.HandleFunc("/OrderServlet/findorderByopenid", model.OrderGetAll)
	http.HandleFunc("/OrderServlet/findOrdersByStatus", model.OrderGetAll)
	http.HandleFunc("/OrderServlet/addOrder", model.OrderAdd)
	http.HandleFunc("/OrderServlet/payOrder", model.OrderPay)

	http.HandleFunc("/AddressServlet/findAll", model.UserGetAddresses)
	http.HandleFunc("/AddressServlet/addAddress", model.UserAddAddress)
	http.HandleFunc("/AddressServlet/remove", model.UserRemoveAddress)

	http.Handle("/pub/", http.StripPrefix("/pub/", http.FileServer(http.Dir("./pub"))))

	e := http.ListenAndServe(":8080", nil)
	if e != nil {
		fmt.Println(e)
		return
	}
}
