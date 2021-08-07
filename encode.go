package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"mime/quotedprintable"
)

// encode bytes quoted by ENCODING t
func EncodeQuoted(s []byte, t string) (string, error) {
	// =?{charset}?{encoding}?{encoded_text}?=
	var encodedTxt string
	var encoding string
	switch t {
	case EncodingB:
		encoding = t
		encodedTxt = base64.StdEncoding.EncodeToString(s)
		break
	case EncodingQ:
		encoding = t
		buf := new(bytes.Buffer)
		w := quotedprintable.NewWriter(buf)

		if _, err := w.Write(s); err != nil {
			return "", err
		}
		if err := w.Close(); err != nil {
			return "", err
		}
		encodedTxt = buf.String()
		break
	default:
		return "", errors.New("encoding type NOT in [B|Q]")
	}

	return fmt.Sprintf("=?UTF-8?%s?%s?=", encoding, encodedTxt), nil
}
