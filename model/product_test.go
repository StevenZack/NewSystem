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
	ret := addPro("https://img.alicdn.com/imgextra/https://img.alicdn.com/imgextra/i2/2265724160/TB2gCQTXpyZBuNjt_jJXXbDlXXa_!!2265724160.jpg_430x430q90.jpg", "短袖t恤女夏装2018新款宽松韩范学生条纹半袖女装上衣服", "女装", 59)
	ret = addPro("https://img.alicdn.com/imgextra/https://img.alicdn.com/imgextra/i2/2265724160/TB2wj4gr79WBuNjSspeXXaz5VXa_!!2265724160.jpg_430x430q90.jpg", "夏装短袖女2018新款白色t恤圆领宽松半袖纯棉体恤女装衣服", "女装", 59)
	t.Log(ret)
}
func addPro(is interface{}, name, typ string, price int) string {
	var images []string
	if v, ok := is.(string); ok {
		images = append(images, v)
	} else if v, ok := is.([]string); ok {
		images = v
	}
	b, e := json.Marshal(Product{
		Name:   name,
		Type:   typ,
		Images: images,
		Price:  price,
	})
	if e != nil {
		return e.Error()
	}
	rp, e := http.Post("http://101.200.54.63:8080/ProServlet/add", "application/json", bytes.NewReader(b))
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
