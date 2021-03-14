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

	bpEnd := 4

	var byteSlice []byte
	for ; len(dst2q) != 0 ; {
		// 1文字目のbit幅を取得するためunicodeのビット幅の情報を持つ部分を取得
		// 最初の4bitに情報が入ってきており、ここでは4進数に変換されているので4桁を取ってくれば良い
		bitWidthFour := dst2q[:bpEnd]

		firstCharBitWidth, err := getOctetsFromBitPattern(bitWidthFour)
		if err != nil {
			return "", err
		}

		bpEnd = firstCharBitWidth * 4
		// 1文字目のbyte数がわかったのでbyte数*4の文だけsubstringを取得してくる
		// (binaryのデータでは 1110 1001 なのが 3221 となるため)
		charBytes := string([]rune(dst2q)[:bpEnd])
		dst2q = string([]rune(dst2q)[bpEnd:])

		// ここで得られた値を直接変えてしまうとだめなので１文字ずつ変換してあげる必要がある

		bs, err := decodeCharFromQuadrant(charBytes)
		if err != nil {
			return "", err
		}

		byteSlice = append(byteSlice, bs...)
	}

	return string(byteSlice), nil
}

// 文字のBit幅を取得すためにUnicodeのビットパターンから有効ビット数を取得する
// 最初の1byteデータを入力して頭についている1の数を数え上げる
func getOctetsFromBitPattern(firstByte string) (int, error) {
	// 4進数からUint8に変換する
	bwUint, err := strconv.ParseUint(firstByte, 4, 64)
	if err != nil {
		return -1, err
	}

	// 必要なbyte数を数える
	var bitWidth int
	for i := 8; i >= 0; i-- {
		if refBitFromByte(int(bwUint), i-1) == 0 {
			break
		}
		bitWidth++
	}

	return bitWidth, nil
}

// 4進数で入ってきたデータをもとに文字を生成する
// 4進数のデータを1つずつ2桁の2進数に変換し結合、それをbyteデータに変換して返す
func decodeCharFromQuadrant(charBytes string) ([]byte, error) {
	var s string
	var byteSlice []byte
	for i, fc := range strings.Split(charBytes, "") {
		s2i, err := strconv.Atoi(fc)
		if err != nil {
			return byteSlice, err
		}
		s += fmt.Sprintf("%02b", s2i)

		// 4進数のデータなので1byteは4文字に相当する
		// 1byteが処理し終わったあとにbyte型に変換してbyteSliceに追加する
		if (i+1) % 4 == 0 {
			bi, err := strconv.ParseUint(s, 2, 32)
			if err != nil {
				return byteSlice, err
			}
			byteSlice = append(byteSlice, byte(bi))
			s = ""
		}
	}
	return byteSlice, nil
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
