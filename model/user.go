package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"net/http"
	"os"
)

type User struct {
	Base
	OpenId    string   `json:"openId"`
	Addresses []string `json:"addresses"`
	Nickname  string   `json:"nickName"`
	Password  string   `json:"password"`
	Avatar    string   `json:"avatar"`
}

func UserRegister(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	r.ParseMultipartForm(32 << 20)
	u := User{}
	u.Nickname = r.FormValue("nickName")
	u.Password = r.FormValue("password")
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	c := s.DB(db).C(cuser)
	count, e := c.Find(bson.M{"nickname": u.Nickname}).Count()
	if e != nil {
		returnErr(w, e)
		return
	}
	if count > 0 {
		returnErr(w, "该昵称已被注册")
		return
	}
	u.OpenId = newToken()
	f, header, e := r.FormFile("file")
	if e != nil {
		returnErr(w, e)
		return
	}
	defer f.Close()
	foDir := pubDir + "avatar/"
	e = os.MkdirAll(foDir, 0755)
	if e != nil {
		returnErr(w, e)
		return
	}
	foPath := foDir + u.OpenId + getFormatFromFileName(header.Filename)
	fo, e := os.Create(foPath)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer fo.Close()
	_, e = io.Copy(fo, f)
	if e != nil {
		returnErr(w, e)
		return
	}
	u.Avatar = "/" + foPath
	e = c.Insert(u)
	if e != nil {
		returnErr(w, e)
		return
	}
	returnData(w, u)
}
func UserLogin(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	r.ParseMultipartForm(20 << 10)
	u := User{}
	u.Nickname = r.FormValue("nickName")
	u.Password = r.FormValue("password")
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	c := s.DB(db).C(cuser)
	gu := User{}
	e = c.Find(bson.M{"nickname": u.Nickname}).One(&gu)
	if e != nil {
		returnErr(w, e)
		return
	}
	if gu.Password != u.Password {
		returnErr(w, "密码不正确")
		return
	}
	returnData(w, gu)
}
func UserAddAddress(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	r.ParseMultipartForm(10 << 10)
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
	e = c.Insert(gu)
	if e != nil {
		returnErr(w, e)
		return
	}
	returnData(w, Base{Status: "OK"})
}
func UserGetAddresses(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	r.ParseMultipartForm(10 << 10)
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
