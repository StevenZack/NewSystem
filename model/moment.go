package model

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

type Moment struct {
	Base
	Text       string   `json:"text"`
	Address    string   `json:"address"`
	Images     []string `json:"images"`
	OpenId     string   `json:"openId"`
	CreateTime string   `json:"createTime"`
}

func MomentUpload(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	r.ParseMultipartForm(32 << 20)
	// text := r.FormValue("text")
	// address := r.FormValue("address")
	openId := r.FormValue("openId")
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	cu := s.DB(db).C(cuser)
	// cm := s.DB(db).C(cmoment)
	count, e := cu.Find(bson.M{"openid": openId}).Count()
	if e != nil {
		returnErr(w, e)
		return
	}
	if count < 1 {
		returnErr(w, "用户不存在，请先注册")
		return
	}
}
func MomentQuery(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	fmt.Println(len(r.MultipartForm.Value["text"]))
	fmt.Fprint(w, "OK")
}
