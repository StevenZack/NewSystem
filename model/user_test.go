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
