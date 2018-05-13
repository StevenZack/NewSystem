package model

import (
	"net/http"
)

type Moment struct {
	Base
	Text    string   `json:"text"`
	Address string   `json:"address"`
	Images  []string `json:"images"`
	OpenId  string   `json:"openId"`
}

func MomentUpload(w http.ResponseWriter, r *http.Request) {

}
func MomentQuery(w http.ResponseWriter, r *http.Request) {

}
