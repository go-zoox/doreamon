package qrcode

import q "github.com/skip2/go-qrcode"

func Encode(text string) []byte {
	png, err := q.Encode(text, q.Medium, 256)
	if err != nil {
		return []byte{}
	}

	return png
}
