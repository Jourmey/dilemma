package youget

import (
	"testing"
)

func TestYouGet_Info(t *testing.T) {
	y := NewYouGet(nil)
	url := "https://www.bilibili.com/video/BV1iy4y1G7b8"

	res, err := y.Info(url)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestYouGet_Download(t *testing.T) {
	y := NewYouGet(nil)
	url := "https://www.bilibili.com/video/BV1iy4y1G7b8"

	res, err := y.Download(url, "dash-flv360", "./video")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
