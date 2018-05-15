package model

import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Base
	ProductId   string   `json:"productId"`
	Images      []string `json:"images"`
	Type        string   `json:"type"`
	Price       int      `json:"price"`
	Name        string   `json:"name"`
	EnglishName string   `json:"englishName"`

	Names string `json:"names"`
}

func ProFindByName(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	query := r.FormValue("query")
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	cp := s.DB(db).C(cproduct)
	var backData struct {
		Base
		Products []Product `json:"products"`
	}
	e = cp.Find(bson.M{"$text": bson.M{"$search": splitHans(query)}}).All(&backData.Products)
	if e != nil {
		returnErr(w, e)
		return
	}
	backData.Status = "OK"
	returnData(w, backData)
}
func ProFindByEN(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	query := r.FormValue("query")
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	cp := s.DB(db).C(cproduct)
	var backData struct {
		Base
		Products []Product `json:"products"`
	}
	e = cp.Find(bson.M{"$text": bson.M{"$search": splitHans(query)}}).All(&backData.Products)
	if e != nil {
		returnErr(w, e)
		return
	}
	backData.Status = "OK"
	returnData(w, backData)
}
func ProFindByType(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	query := r.FormValue("query")
	typ := r.FormValue("type")
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	cp := s.DB(db).C(cproduct)
	var backData struct {
		Base
		Products []Product `json:"products"`
	}
	e = cp.Find(bson.M{"$text": bson.M{"$search": splitHans(query)}, "type": typ}).All(&backData.Products)
	if e != nil {
		returnErr(w, e)
		return
	}
	backData.Status = "OK"
	returnData(w, backData)
}
func ProFuzzyFind(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	query := r.FormValue("query")
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	cp := s.DB(db).C(cproduct)
	var backData struct {
		Base
		Products []Product `json:"products"`
	}
	e = cp.Find(bson.M{"$text": bson.M{"$search": splitHans(query)}}).All(&backData.Products)
	if e != nil {
		returnErr(w, e)
		return
	}
	backData.Status = "OK"
	returnData(w, backData)
}
func ProAdd(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	b, e := ioutil.ReadAll(r.Body)
	if e != nil {
		returnErr(w, e)
		return
	}
	pro := Product{}
	e = json.Unmarshal(b, &pro)
	if e != nil {
		returnErr(w, e)
		return
	}
	pro.ProductId = newToken()
	pro.Names = splitHans(pro.Name) + " " + splitHans(pro.EnglishName)
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	e = s.DB(db).C(cproduct).Insert(pro)
	if e != nil {
		returnErr(w, e)
		return
	}
	returnData(w, Base{Status: "OK"})
}
