package main

import "fmt"

type IEncoder interface {
	Encode(value string)
	Decode(value string)
}

type XEncoder struct {
}

type YEncoder struct {
}

func (XEncoder *XEncoder) Encode(value string) {
	fmt.Println("XEncoder ile encode edildi.")
}

func (XEncoder *XEncoder) Decode(value string) {
	fmt.Println("XEncoder ile decode edildi.")
}

func (YEncoder *YEncoder) Encode(value string) {
	fmt.Println("YEncoder ile encode edildi.")
}

func (YEncoder *YEncoder) Decode(value string) {
	fmt.Println("YEncoder ile decode edildi.")
}

func main() {
	var encoder IEncoder

	encoder = &XEncoder{}

	encoder.Encode("qweqwe")
}
