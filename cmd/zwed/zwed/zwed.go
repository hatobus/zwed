package zwed

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

var (
	ZWSPS map[string]string = map[string]string{
		"0": "\u200b",
		"1": "\u200c",
		"2": "\u200d",
		"3": "\ufeff",
	}
	DELIM = "\u2028"
)

func decimalToQuat(num int64) string {
	var mod int64
	// 底の変換公式より
	// log_4(N) = log_2(N) / 2
	var strarr = make([]int64, int(math.Log2(float64(num))/2))
	for true {
		mod = num % 4
		strarr = append(strarr, mod)
		num = num / 4
		if num == 0 {
			break
		}
	}

	log.Println(strarr)

	var outbuf string
	for i := len(strarr) - 1; i >= 0; i-- {
		outbuf += string(strarr[i])
	}

	return outbuf
}

func Decode(dst string) (string, error) {
	dst = strings.TrimRight(dst, "\n")

	runeslice := []rune(dst)

	var cs string
	var ci int64
	var quatarr = make([]string, len(dst))
	for _, c := range runeslice {
		cs = fmt.Sprintf("%d", c)
		ci, _ = strconv.ParseInt(cs, 10, 64)
		quatarr = append(quatarr, decimalToQuat(ci))
	}

	var out string
	for _, qs := range quatarr {
		for digit, zwc := range ZWSPS {
			qs = strings.Replace(qs, digit, zwc, -1)
		}
		out += qs + DELIM
	}

	return out, nil
}
