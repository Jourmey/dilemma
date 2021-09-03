package main

import (
	"dilemma/youget"
	"fmt"
)

func main() {
	y := youget.NewYouGet()
	url := "https://www.bilibili.com/video/BV1iy4y1G7b8"

	res, err := y.Info(url)
	fmt.Println(res, err)
}
