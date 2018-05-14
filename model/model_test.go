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
func Test_splitHans(t *testing.T) {
	t.Log(splitHans("asd汉子 a阿萨德 asd qwA有"))
}
