package model

import (
	"os"
	"testing"
)

func Test_MomentQuery(t *testing.T) {
	fi, e := os.Stat("/home/asd/a.txt")
	if e != nil {
		t.Log(e)
		return
	}
	t.Log(fi.Size())
}
