package zwed

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	ZWSPS map[string]string = map[string]string{
		"0": "\u200b",
		"1": "\u200c",
		"2": "\u200d",
		"3": "\u200e",
	}
	DELIM = "\u034f"
)

func SetZWSpace(m map[string]string) {
	ZWSPS = m
}

func GetZWSpace() map[string]string {
	return ZWSPS
}

func SetDELIM(delim string) {
	DELIM = delim
}

func GetDELIM() string {
	return DELIM
}

func decimalToQuat(num int64) string {
	var mod int64

	// 底の変換公式より
	// log_4(N) = log_2(N) / 2
	var strarr = make([]int64, 0, int(math.Log2(float64(num))/2)+1)

	for true {
		mod = num % 4
		strarr = append(strarr, mod)
		// int同士の割り算はintになる
		num = num / 4
		if num == 0 {
			break
		}
	}

	var outbuf int64 = 0
	for i := len(strarr) - 1; i >= 0; i-- {
		outbuf = outbuf*10 + strarr[i]
	}

	return strconv.Itoa(int(outbuf))
}

func quatToDecimal(num string) int64 {
	digits := strings.Split(num, "")

	ex := 0

	var out int64
	out = 0
	for i := len(digits) - 1; i >= 0; i-- {
		n, err := strconv.ParseInt(digits[i], 10, 64)
		if err != nil {
			return 0
		}
		out += int64(math.Pow(4, float64(ex))) * n
		ex++
	}

	return out
}

func Encode(dst string) (string, error) {
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

	return out, nil
}

func zwto4(char string) (string, error) {
	for digit, zwc := range ZWSPS {
		char = strings.Replace(char, zwc, digit, -1)
	}

	return char, nil
}

func Decode(dst string) (string, error) {
	dst = strings.TrimRight(dst, "\n")

	chars := strings.Split(dst, DELIM)

	charsdec := make([]int64, 0, len(chars))

	for _, c := range chars {
		quat, err := zwto4(c)
		if err != nil {
			return "", err
		}
		// log.Println(quat)
		dec := quatToDecimal(quat)
		charsdec = append(charsdec, dec)
	}

	var outstr string
	for _, cd := range charsdec {
		outstr += string(cd)
	}

	return outstr, nil
}
