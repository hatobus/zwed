package zwed

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestEncodeAndDecodeString(t *testing.T) {
	testCases := map[string]string{
		"1byte文字の入力": "abcdefg",
		"2byte文字の入力": "¥©",
		"3byte文字の入力": "こんにちは",
		"4byte文字の入力": "𠮷野家で𩸽",

		"複数種類のbyte数の文字が存在する場合": "𠮷野家で𩸽を食べたhatobus ¥1000",
	}

	for testName, tc := range testCases {
		tc := tc
		t.Run(testName, func(t *testing.T) {
			zws, err := Encode(tc)
			if err != nil {
				t.Fatal(err)
			}

			dec, err := Decode(zws)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(dec, tc); diff != "" {
				t.Fatalf("unexpected output, diff = %v\n", diff)
			}
		})
	}
}
