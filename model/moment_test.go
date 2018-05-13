package model

import (
	"testing"
)

func Test_MomentQuery(t *testing.T) {
	m := make(map[string]interface{})
	m["text"] = "one"
	Post("http://127.0.0.1:8080/MomentServlet/findAll", m)
}
