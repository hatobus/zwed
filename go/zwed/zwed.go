package zwed

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	zwssps = map[string]string{
		"0": "\u200b",
		"1": "\u200c",
		"2": "\u200d",
		"3": "\u200e",
	}
	//delim = "\u034f"
)

func SetZWSpace(m map[string]string) {
	zwssps = m
}

func GetZWSpace() map[string]string {
	return zwssps
}

func Encode(dst string) (string, error) {
	byteSlice := []byte(dst)

	var quatarr = make([]string, 0)

	// 4進数に変換する
	for _, bs := range byteSlice {
			quatarr = append(quatarr, decimalToQuat(int64(bs)))
	}

	var qs string
	for _, quat := range quatarr {
		for digit, zwc := range zwssps {
			quat = strings.Replace(quat, digit, zwc, -1)
		}
		qs += quat
	}

	return qs, nil
}

func Decode(dst string) (string, error) {
	dst = strings.TrimRight(dst, "\n")

	dst2q, err := zwto4(dst)
	if err != nil {
		return "", err
	}

	bpStart := 0
	bpEnd := 4

	// 1文字目のbit幅を取得するためunicodeのビット幅の情報を持つ部分を取得
	// 最初の4bitに情報が入ってきており、ここでは4進数に変換されているので4桁を取ってくれば良い
	bitWidthFour := dst2q[:bpEnd]

	// 4進数からUint8に変換する
	bwUint, err := strconv.ParseUint(bitWidthFour, 4, 8)
	if err != nil {
		return "", err
	}

	// 必要なbyte数を数える
	var firstCharBitWidth int
	for i := 8; i >= 0; i-- {
		if refBitFromByte(int(bwUint), i-1) == 0 {
			break
		}
		firstCharBitWidth++
	}

	// 1文字目のbyte数がわかったのでbyte数*4の文だけsubstringを取得してくる
	// (binaryのデータでは 1110 1001 なのが 3221 となるため)
	firstCharBytes := string([]rune(dst2q)[bpStart:(firstCharBitWidth*4)])

	// ここで得られた値を直接変えてしまうとだめなので１文字ずつ変換してあげる必要がある
	byteSlice := make([]byte, 0, 0)
	var s string
	for i, fc := range strings.Split(firstCharBytes, "") {
		s2i, err := strconv.Atoi(fc)
		if err != nil {
			return "", err
		}
		s += fmt.Sprintf("%02b", s2i)
		if (i+1) % 4 == 0 {
			bi, err := strconv.ParseUint(s, 2, 32)
			if err != nil {
				return "", err
			}
			byteSlice = append(byteSlice, byte(bi))
			s = ""
		}
	}

	return string(byteSlice), nil
}

func refBitFromByte(b int, i int) int {
	return (b >> i) & 1
}

func zwto4(char string) (string, error) {
	for digit, zwc := range zwssps {
		char = strings.Replace(char, zwc, digit, -1)
	}

	return char, nil
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

