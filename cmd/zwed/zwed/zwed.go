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
	var strarr = make([]int64, 0, int(math.Log2(float64(num))/2)+1)

	for true {
		mod = num % 4
		strarr = append(strarr, mod)
		num = num / 4
		if num == 0 {
			break
		}
	}

	log.Println(strarr)

	var outbuf int64 = 0
	for i := len(strarr) - 1; i >= 0; i-- {
		outbuf = outbuf*10 + strarr[i]
	}

	return strconv.Itoa(int(outbuf))
}

func quatToDecimal(num int64) int64 {
	return 0
}

func Decode(dst string) (string, error) {
	dst = strings.TrimRight(dst, "\n")

	runeslice := []rune(dst)

	var cs string
	var ci int64
	var quatarr = make([]string, 0)

	// 4進数に変換する
	for _, c := range runeslice {
		cs = fmt.Sprintf("%d", c)
		ci, _ = strconv.ParseInt(cs, 10, 64)
		quatarr = append(quatarr, decimalToQuat(ci))
	}

	var out string
	// 変換された文字をDELIMを挟めて文字列へ
	for _, qs := range quatarr {
		for digit, zwc := range ZWSPS {
			qs = strings.Replace(qs, digit, zwc, -1)
		}
		out += qs + DELIM
	}

	// 最後のDELIMはいらないのでtrim
	out = strings.TrimRight(out, DELIM)
	log.Println(len(out))

	return out, nil
}
