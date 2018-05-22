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
	ret := addPro("https://img.alicdn.com/imgextra/i3/TB1dTSafL6H8KJjy0Fjw8mXepXa_013809.jpg_430x430q90.jpg", "夏季减肥修正 左旋肉碱茶多酚片 0.8g/片*60片男女减肥", "保健品", 68)
	ret = addPro("https://img.alicdn.com/imgextra/i2/TB1.UKKNpXXXXcnXFXXXXXXXXXX_!!0-item_pic.jpg_430x430q90.jpg", "CENTRUM/善存R佳维片 1.33g/片*150片+133.g/片*60片复合维生素", "保健品", 189)
	ret = addPro("https://img.alicdn.com/imgextra/i4/2549841410/TB2pdGsr1uSBuNjSsplXXbe8pXa_!!2549841410-0-sm.jpg_430x430q90.jpg", "【直营】swisse女士复合维生素片120片 女士保健品", "保健品", 249)
	ret = addPro("https://img.alicdn.com/imgextra/i7/TB1VAoVPVXXXXcjXFXXy74f.VXX_110843.jpg_430x430q90.jpg", "脑白金口服液10天剂量礼盒装改善深睡眠更年期保健品送礼", "保健品", 128)
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
