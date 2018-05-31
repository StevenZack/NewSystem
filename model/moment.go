package model

import (
	"fmt"
	"github.com/StevenZack/tools"
	"io"
	"net/http"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
	openId := handleStr(r, "openId")
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	cu := s.DB(db).C(cuser)
	cm := s.DB(db).C(cmoment)
	count, e := cu.Find(bson.M{"openid": openId}).Count()
	if e != nil {
		returnErr(w, e)
		return
	}
	if count < 1 {
		returnErr(w, "用户不存在，请先注册")
		return
	}
	text := handleStr(r, "text")
	var images []string
	fis := r.MultipartForm.File["images"]
	for _, v := range fis {
		fi, e := v.Open()
		if e != nil {
			fmt.Fprint(w, e)
			return
		}
		os.MkdirAll(pubDir+"images/", 0755)
		rpath := pubDir + "images/" + tools.NewToken() + getFormatFromFileName(v.Filename)
		_, e = os.Stat(rpath)
		if e == nil {
			rpath = pubDir + "images/" + tools.NewNumToken() + getFormatFromFileName(v.Filename)
		}
		fo, e := os.OpenFile(rpath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if e != nil {
			fmt.Fprint(w, e)
			return
		}
		io.Copy(fo, fi)
		fmt.Println(rpath)
		images = append(images, host+rpath)
		fo.Close()
	}
	e = cm.Insert(Moment{Text: text, Images: images, OpenId: openId, CreateTime: tools.GetTimeStrNow()})
	if e != nil {
		returnErr(w, e)
		return
	}
	returnData(w, Base{Status: "OK"})
}
func MomentQuery(w http.ResponseWriter, r *http.Request) {
	handleCon(w)
	openid := r.FormValue("openId")
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		returnErr(w, e)
		return
	}
	defer s.Close()
	count, e := s.DB(db).C(cuser).Find(bson.M{"openid": openid}).Count()
	if e != nil {
		returnErr(w, e)
		return
	}
	if count < 1 {
		returnErr(w, "账号不存在，请重新登录")
		return
	}
	var backData struct {
		Base
		Moments []Moment `json:"moments"`
	}
	e = s.DB(db).C(cmoment).Find(bson.M{"openid": openid}).All(&backData.Moments)
	if e != nil {
		returnErr(w, e)
		return
	}
	backData.Status = "OK"
	returnData(w, backData)
}
