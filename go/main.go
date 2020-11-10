package main

import (
	"fmt"

	"github.com/hatobus/zwed/go/zwed"
)

func main() {
	test := "こんにちは世界"
	fmt.Printf("Original message: %v\n", test)
	zws, _ := zwed.Encode(test)
	fmt.Printf("zws :%v, len : %v \n", zws, len(zws))
	dec, _ := zwed.Decode(zws)
	fmt.Printf("Decoded string : %v \n", dec)
}
