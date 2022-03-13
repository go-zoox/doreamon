package aes

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
)

type Encoding interface {
	EncodeToString(src []byte) string
	DecodeString(s string) ([]byte, error)
}

type HexEncoding struct{}

func (e *HexEncoding) EncodeToString(src []byte) string {
	return hex.EncodeToString(src)
}

func (e *HexEncoding) DecodeString(s string) ([]byte, error) {
	return hex.DecodeString(s)
}

type Base64Encoding struct{}

func (e *Base64Encoding) EncodeToString(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func (e *Base64Encoding) DecodeString(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

type Base32Encoding struct{}

func (e *Base32Encoding) EncodeToString(src []byte) string {
	return base32.StdEncoding.EncodeToString(src)
}

func (e *Base32Encoding) DecodeString(s string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(s)
}
