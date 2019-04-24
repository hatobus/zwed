package main

import (
	"fmt"

	"github.com/hatobus/zwed/cmd/zwed/zwed"
)

func main() {
	test := "ほげほげふがふが"
	zws, _ := zwed.Encode(test)
	fmt.Printf("zws :%v, len : %v \n", zws, len(zws))
	dec, _ := zwed.Decode(zws)
	fmt.Printf("Decoded string : %v \n", dec)
}
