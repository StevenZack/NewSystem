package model

import (
	"encoding/json"
	"fmt"
	"github.com/StevenZack/tools"
	"net/http"
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

func getFormatFromFileName(fname string) string {
	for i := len(fname) - 1; i > -1; i-- {
		if fname[i:i+1] == "." {
			return fname[i:]
		}
	}
	return ""
}

func handleV(r *http.Request, key string) string {
	vs := r.MultipartForm.File[key]
	if len(vs) > 0 {
		fi, e := vs[0].Open()
		if e != nil {
			fmt.Println(e)
			return ""
		}
		str := tools.ReadAll(fi)
		fmt.Println(str)
		fi.Close()
		return str
	}
	return ""
}
