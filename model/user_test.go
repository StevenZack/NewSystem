package model

import (
	"bytes"
	"io"
	"io/ioutil"
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
func Post(url string, m map[string]interface{}) (string, error) {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	for k, v := range m {
		if vv, ok := v.(string); ok {
			str, e := w.CreateFormField(k)
			if e != nil {
				continue
			}
			str.Write([]byte(vv))
			continue
		}
		if vv, ok := v.(*os.File); ok {
			fo, e := w.CreateFormFile(k, vv.Name())
			if e != nil {
				continue
			}
			io.Copy(fo, vv)
			vv.Close()
			continue
		}
	}
	w.Close()
	r, e := http.NewRequest("POST", url, buf)
	if e != nil {
		return "", e
	}
	var client http.Client
	rp, e := client.Do(r)
	if e != nil {
		return "", e
	}
	defer rp.Body.Close()
	b, e := ioutil.ReadAll(rp.Body)
	return string(b), e
}