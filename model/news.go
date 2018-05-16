package model

import (
	"encoding/json"
	"github.com/StevenZack/tools"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	handleCon(w)
	typ := r.FormValue("type")
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	cn := s.DB(db).C(cnews)
	var backData struct {
		Base
		NewsArray []News `json:"newsArrary"`
	}
	e = cn.Find(bson.M{"type": typ}).All(&backData.NewsArray)
	if e != nil {
		returnErr(w, e)
		return
	}
	backData.Status = "OK"
	returnData(w, backData)
}
func NewsGetDetail(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	news_id := r.FormValue("news_id")
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	cn := s.DB(db).C(cnews)
	gn := News{}
	e = cn.Find(bson.M{"news_id": news_id}).One(&gn)
	if e != nil {
		returnErr(w, e)
		return
	}
	gn.Status = "OK"
	returnData(w, gn)
}
func NewsFindAll(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	var backData struct {
		Base
		NewsArray []News `json:"newsArray"`
	}
	e = s.DB(db).C(cnews).Find(nil).All(&backData.NewsArray)
	if e != nil {
		returnErr(w, e)
		return
	}
	backData.Status = "OK"
	returnData(w, backData)
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
	n.News_id = tools.NewToken()
	n.CreateTime = tools.GetTimeStrNow()
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
