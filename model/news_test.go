package model

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func Test_NewsAdd(t *testing.T) {
}
func addNews(title, content, publisher, typ string, images []string) string {
	n := News{Title: title, Content: content, Publisher_title: publisher, Type: typ, Images: images}
	b, e := json.Marshal(n)
	if e != nil {
		return e.Error()
	}
	br := bytes.NewReader(b)
	rp, e := http.Post("http://127.0.0.1:8080/NewsServlet/add", "application/json", br)
	if e != nil {
		return e.Error()
	}
	return readAll(rp.Body)
}
