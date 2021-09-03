package youget

import (
	"testing"
)

func TestYouGet_Info(t *testing.T) {
	y := NewYouGet()
	url := "https://www.bilibili.com/video/BV1iy4y1G7b8"

	res, err := y.Info(url)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
