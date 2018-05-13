package model

import (
	"net/http"
)

type Product struct {
	Base
	ProductId   string   `json:"productId"`
	Images      []string `json:"images"`
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	EnglishName string   `json:"englishName"`
	Names       string
}

func ProFindByName(w http.ResponseWriter, r *http.Request) {
}
func ProFindByEN(w http.ResponseWriter, r *http.Request) {

}
func ProFindByType(w http.ResponseWriter, r *http.Request) {

}
func ProFuzzyFind(w http.ResponseWriter, r *http.Request) {

}
