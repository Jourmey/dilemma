package main

import (
	"dilemma/youget"
	"fmt"
)

func main() {
	y := youget.NewYouGet()
	url := "https://www.bilibili.com/video/BV1iy4y1G7b8"

	res, err := y.Download(url, "dash-flv360", "./video")
	fmt.Println(res, err)
}
