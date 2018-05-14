package model

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Base struct {
	Status     string `json:"status"`
	StatusInfo string `json:"statusInfo"`
}

var (
	pubDir   string = "pub/"
	mongoDB  string = "127.0.0.1"
	host     string = "http://101.200.54.63:8080/"
	db       string = "NewSystem"
	cuser    string = "user"
	corder   string = "order"
	cmoment  string = "moment"
	cnews    string = "news"
	cproduct string = "product"
)

func handleCon(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}
func returnData(w http.ResponseWriter, i interface{}) {
	b, e := json.Marshal(i)
	if e != nil {
		fmt.Println(e)
		returnData(w, Base{Status: "ERR", StatusInfo: e.Error()})
		return
	}
	w.Write(b)
}
func returnErr(w http.ResponseWriter, e interface{}) {
	if v, ok := e.(error); ok {
		returnData(w, Base{Status: "ERR", StatusInfo: v.Error()})
		return
	}
	if v, ok := e.(string); ok {
		returnData(w, Base{Status: "ERR", StatusInfo: v})
		return
	}
}
func newToken() string {
	ct := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(ct, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}
func getFormatFromFileName(fname string) string {
	for i := len(fname) - 1; i > -1; i-- {
		if fname[i:i+1] == "." {
			return fname[i:]
		}
	}
	return ""
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
func getTimeNow() string {
	return time.Now().Format("2006-01-01 15:03:02")
}
func handleV(r *http.Request, key string) string {
	vs := r.MultipartForm.File[key]
	if len(vs) > 0 {
		fi, e := vs[0].Open()
		if e != nil {
			fmt.Println(e)
			return ""
		}
		str := readAll(fi)
		fmt.Println(str)
		fi.Close()
		return str
	}
	return ""
}
func readAll(r io.Reader) string {
	b, e := ioutil.ReadAll(r)
	if e != nil {
		return ""
	}
	return string(b)
}
