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
	ret := addPro("玻尿酸密集补水面膜 修护保湿滋润", "Home Facial Pro", "面膜", 99, "https://g-search1.alicdn.com/img/bao/uploaded/i4/imgextra/i2/117606064/TB2Awi4gGSWBuNjSsrbXXa0mVXa_!!0-saturn_solar.jpg_180x180.jpg_.webp")
	ret = addPro("云南白药采之汲自然原酵系列天然面膜深度补水保湿", "", "面膜", 59,
		"https://g-search3.alicdn.com/img/bao/uploaded/i4/i4/2077111000/TB2sHU.rGSWBuNjSsrbXXa0mVXa_!!2077111000-0-item_pic.jpg_180x180.jpg_.webp")
	ret = addPro("珀莱雅人鱼公主玻尿酸面膜贴女海藻补水保湿提亮清", "", "面膜", 99,
		"https://g-search3.alicdn.com/img/bao/uploaded/i4/i4/379424083/TB2qGQ4rhWYBuNjy1zkXXXGGpXa_!!379424083-0-item_pic.jpg_180x180.jpg_.webp")
	ret = addPro("[送13片]杰威尔男士面膜美白补水保湿控油祛痘去黑", "", "面膜", 69,
		"https://g-search3.alicdn.com/img/bao/uploaded/i4/i1/756239978/TB1VMMur_lYBeNjSszcXXbwhFXa_!!0-item_pic.jpg_180x180.jpg_.webp")
	t.Log(ret)
}
func addPro(name, englishName, typ string, price int, is interface{}) string {
	var images []string
	if v, ok := is.(string); ok {
		images = append(images, v)
	} else if v, ok := is.([]string); ok {
		images = v
	}
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
