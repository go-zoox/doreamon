package base64

import (
	"encoding/base64"
)

func Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func Decode(text string) string {
	v, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return ""
	}

	return string(v)
}
