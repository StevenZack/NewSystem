package model

import (
	"testing"
)

func Test_getFormatFromFileName(t *testing.T) {
	t.Log(getFormatFromFileName("asd.png"))
}
func Test_getTimeNow(t *testing.T) {
	t.Log(getTimeNow())
}
