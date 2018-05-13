package model

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
