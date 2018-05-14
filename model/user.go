package model

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Base
	OpenId    string   `json:"openId"`
	Addresses []string `json:"addresses"`
	Nickname  string   `json:"nickName"`
	Avatar    string   `json:"avatar"`
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	openId := r.FormValue("openId")
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	c := s.DB(db).C(cuser)
	gu := User{}
	count, e := c.Find(bson.M{"openid": openId}).Count()
	if e != nil {
		returnErr(w, e)
		return
	}
	if count > 0 { //login
		e = c.Find(bson.M{"openid": openId}).One(&gu)
		if e != nil {
			returnErr(w, e)
			return
		}
		returnData(w, Base{Status: "OK"})
	}
	//register
	gu.OpenId = openId
	gu.Nickname = r.FormValue("nickName")
	gu.Avatar = r.FormValue("head")
	e = c.Insert(gu)
	if e != nil {
		returnErr(w, e)
		return
	}
	returnData(w, Base{Status: "OK"})
}
func UserAddAddress(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	openid := r.FormValue("openId")
	addr := r.FormValue("address")
	gu := User{}
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	c := s.DB(db).C(cuser)
	e = c.Find(bson.M{"openid": openid}).One(&gu)
	if e != nil {
		returnErr(w, e)
		return
	}
	gu.Addresses = append(gu.Addresses, addr)
	e = c.Update(bson.M{"openid": openid}, gu)
	if e != nil {
		returnErr(w, e)
		return
	}
	returnData(w, Base{Status: "OK"})
}
func UserGetAddresses(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	openid := r.FormValue("openId")
	gu := User{}
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	c := s.DB(db).C(cuser)
	e = c.Find(bson.M{"openid": openid}).One(&gu)
	if e != nil {
		returnErr(w, e)
		return
	}
	returnData(w, gu)
}
