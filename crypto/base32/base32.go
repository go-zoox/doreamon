package base32

import (
	"encoding/base32"
)

func Encode(text string) string {
	return base32.StdEncoding.EncodeToString([]byte(text))
}

func Decode(text string) string {
	v, err := base32.StdEncoding.DecodeString(text)
	if err != nil {
		return ""
	}

	return string(v)
}
