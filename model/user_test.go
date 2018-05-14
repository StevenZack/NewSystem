package model

import (
	"bytes"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
)

func Test_UserRegister(t *testing.T) {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	fw, e := w.CreateFormFile("file", "one.png")
	if e != nil {
		t.Error(e)
		return
	}
	fd, e := os.Open("/home/asd/Pictures/one.png")
	if e != nil {
		t.Error(e)
		return
	}
	defer fd.Close()
	_, e = io.Copy(fw, fd)
	if e != nil {
		t.Error(e)
		return
	}
	nick, e := w.CreateFormField("nickName")
	if e != nil {
		t.Error(e)
		return
	}
	nick.Write([]byte("stevenzack"))
	password, e := w.CreateFormField("password")
	if e != nil {
		t.Error(e)
		return
	}
	password.Write([]byte("123456"))
	w.Close()
	r, e := http.NewRequest("POST", "http://127.0.0.1:8080/UserServlet/register", buf)
	if e != nil {
		t.Error(e)
		return
	}
	r.Header.Set("Content-Type", w.FormDataContentType())
	var client http.Client
	rp, e := client.Do(r)
	if e != nil {
		t.Error(e)
		return
	}
	defer rp.Body.Close()
	io.Copy(os.Stdout, rp.Body)
}
func Test_UserAddAddress(t *testing.T) {
	s, e := mgo.Dial(mongoDB)
	if e != nil {
		return
	}
	defer s.Close()
	gu := User{}
	cu := s.DB(db).C(cuser)
	e = cu.Find(bson.M{"openid": "fe6e1672cadd4f232e81a68ca17f8144"}).One(&gu)
	if e != nil {
		return
	}
	gu.Addresses = nil
	e = cu.Update(bson.M{"openid": gu.OpenId}, gu)
	if e != nil {
		return
	}
	t.Log("OK")
}
