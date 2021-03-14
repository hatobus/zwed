package main

import (
	"fmt"

	"github.com/hatobus/zwed/go/zwed"
)

func main() {
	test := "こんにちは世界"
	fmt.Printf("Original message: %v\n", test)
	zws, err := zwed.Encode(test)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("zws :%v, len : %v \n", zws, len(zws))
	dec, err := zwed.Decode(zws)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Decoded string : %v \n", dec)
}
