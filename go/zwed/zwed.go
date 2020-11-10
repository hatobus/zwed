package zwed

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	zwssps map[string]string = map[string]string{
		"0": "\u200b",
		"1": "\u200c",
		"2": "\u200d",
		"3": "\u200e",
	}
	delim = "\u034f"
)

func SetZWSpace(m map[string]string) {
	zwssps = m
}

func GetZWSpace() map[string]string {
	return zwssps
}

func Setdelim(d string) {
	delim = d
}

func Getdelim() string {
	return delim
}

func decimalToQuat(num int64) string {
	var mod int

	// 底の変換公式より
	// log_4(N) = log_2(N) / 2
	var strarr = make([]string, 0, int(math.Log2(float64(num))/2)+1)

	for {
		mod = int(num % 4)
		strarr = append(strarr, strconv.Itoa(mod))
		// int同士の割り算はintになる
		num /= 4
		if num == 0 {
			break
		}
	}

	for i := 0; i < len(strarr)/2; i++ {
		strarr[i], strarr[len(strarr)-i-1] = strarr[len(strarr)-i-1], strarr[i]
	}

	return strings.Join(strarr, "")
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

	qs := make([]string, 0, len(quatarr))

	for _, quat := range quatarr {
		for digit, zwc := range zwssps {
			quat = strings.Replace(quat, digit, zwc, -1)
		}
		qs = append(qs, quat)
	}

	return strings.Join(qs, delim), nil
}

func zwto4(char string) (string, error) {
	for digit, zwc := range zwssps {
		char = strings.Replace(char, zwc, digit, -1)
	}

	return char, nil
}

func Decode(dst string) (string, error) {
	dst = strings.TrimRight(dst, "\n")

	chars := strings.Split(dst, delim)

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
