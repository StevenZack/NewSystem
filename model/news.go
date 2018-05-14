package model

import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"net/http"
)

type News struct {
	Base
	News_id         string   `json:"news_id"`
	Title           string   `json:"title"`
	Content         string   `json:"content"`
	CreateTime      string   `json:"createTime"`
	Publisher_title string   `json:"publisher_title"`
	Images          []string `json:"images"`
	Type            string   `json:"type"`
}

func NewsFindByType(w http.ResponseWriter, r *http.Request) {

}
func NewsGetDetail(w http.ResponseWriter, r *http.Request) {

}
func NewsFindAll(w http.ResponseWriter, r *http.Request) {

}
func NewsAdd(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	b, e := ioutil.ReadAll(r.Body)
	if e != nil {
		returnErr(w, e)
		return
	}
	n := News{}
	e = json.Unmarshal(b, &n)
	if e != nil {
		returnErr(w, e)
		return
	}
	n.News_id = newToken()
	n.CreateTime = getTimeNow()
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	cn := s.DB(db).C(cnews)
	e = cn.Insert(n)
	if e != nil {
		returnErr(w, e)
		return
	}
	returnData(w, Base{Status: "OK"})
}
