package main

import (
	"fmt"
	"testing"
)

func TestQuoted(t *testing.T) {

	var words = []string{
		"测试123文字。",
		"hello man.",
	}

	for _, w := range words {
		encoded, err := EncodeQuoted([]byte(w), EncodingB)
		if err != nil {
			t.Fatal(fmt.Sprintf("%s [%s]",err,w))
		}
		decoded, err := DecodeQuoted(encoded)
		if err != nil {
			t.Fatal(fmt.Sprintf("%s [%s]",err,w))
		}
		if w != decoded {
			t.Fatal(fmt.Sprintf("encode/decode [%s] error", w))
		}
	}

	for _, w := range words {
		encoded, err := EncodeQuoted([]byte(w), EncodingQ)
		if err != nil {
			t.Fatal(fmt.Sprintf("%s [%s]",err,w))
		}
		decoded, err := DecodeQuoted(encoded)
		if err != nil {
			t.Fatal(fmt.Sprintf("%s [%s]",err,w))
		}
		if w != decoded {
			t.Fatal(fmt.Sprintf("encode/decode [%s] error", w))
		}
	}

}
