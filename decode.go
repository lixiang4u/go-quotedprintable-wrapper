package main

import (
	"encoding/base64"
	"errors"
	"io"
	"mime/quotedprintable"
	"regexp"
	"strings"
)

// decode quoted string
func DecodeQuoted(s string) (string, error) {
	var result []byte
	if s == "" {
		return "", nil
	}
	reg, err := regexp.Compile(`(?U)=\?(.+)\?([B|Q])\?(.+)\?=`)
	if err != nil {
		return "", err
	}
	matched := reg.FindAllStringSubmatch(s, -1)
	if len(matched) < 1 {
		return s, errors.New("not parsed from quoted string")
	}
	for _, m := range matched {
		if len(m) < 4 {
			continue
		}
		switch m[2] {
		case EncodingB:
			b, _ := base64.StdEncoding.DecodeString(m[3])
			result = append(result, b...)
			break
		case EncodingQ:
			b, _ := io.ReadAll(quotedprintable.NewReader(strings.NewReader(m[3])))
			result = append(result, b...)
			break
		}
	}
	return string(result), nil
}
