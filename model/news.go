package model

import (
	"net/http"
)

type News struct {
	Base
	News_id         string   `json:"news_id"`
	Title           string   `json:"title"`
	Content         string   `json:"content"`
	CreateTime      string   `json:"createTime"`
	Publisher_title string   `json:"publisher_title"`
	Images          []string `json:"images"`
}

func NewsFindByType(w http.ResponseWriter, r *http.Request) {

}
func NewsGetDetail(w http.ResponseWriter, r *http.Request) {

}
func NewsFindAll(w http.ResponseWriter, r *http.Request) {

}
