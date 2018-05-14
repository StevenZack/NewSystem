package model

import (
	"bytes"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_ProAdd(t *testing.T) {
	ret := addPro("华为P20", "Huawei P20", "手机", 3788, []string{
		"https://img.alicdn.com/bao/uploaded/i6/TB1bh0lq1uSBuNjSsplSWDe8pXa_052653.jpg_b.jpg",
		"https://img.alicdn.com/bao/uploaded/i1/TB1buTWouuSBuNjy1XcCcwYjFXa_095905.jpg_b.jpg",
		"https://img.alicdn.com/bao/uploaded/i5/TB1kJ07ieuSBuNjSsplSOre8pXa_023515.jpg_b.jpg",
	})
	t.Log(ret)
}
func addPro(name, englishName, typ string, price int, images []string) string {
	b, e := json.Marshal(Product{
		Name:        name,
		EnglishName: englishName,
		Type:        typ,
		Images:      images,
		Price:       price,
	})
	if e != nil {
		return e.Error()
	}
	rp, e := http.Post("http://127.0.0.1:8080/ProServlet/add", "application/json", bytes.NewReader(b))
	if e != nil {
		return e.Error()
	}
	data, _ := ioutil.ReadAll(rp.Body)
	return string(data)
}
func Test_ProFindByName(t *testing.T) {
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		t.Error(e)
		return
	}
	defer s.Close()
	cp := s.DB(db).C(cproduct)
	var pros []Product
	e = cp.Find(bson.M{"$text": bson.M{"$search": "华 为"}}).All(&pros)
	if e != nil {
		t.Error(e)
		return
	}
	t.Log(len(pros))
}
