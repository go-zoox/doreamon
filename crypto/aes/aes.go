package aes

import (
	"fmt"
)

var SIZE = []int{128, 192, 256}

func isKeySizeAllow(size int) bool {
	switch size {
	case 128, 192, 256:
		return true
	default:
		return false
	}
}

func isIVSizeAllow(size int) bool {
	switch size {
	case 16:
		return true
	default:
		return false
	}
}

func isNonceSizeAllow(size int) bool {
	switch size {
	case 12:
		return true
	default:
		return false
	}
}

type Aes struct {
	Size     int
	Encoding Encoding
	IV       []byte
}

func (a *Aes) Name() string {
	return fmt.Sprintf("aes-%v-cbc", a.Size)
}

func (a *Aes) BlockSize() int {
	return a.Size / 8
}

func (a *Aes) GetIV(key []byte) []byte {
	if a.IV != nil {
		return a.IV[:16]
	}

	return key[:16]
}

func (a *Aes) GenerateKey() (string, error) {
	return GenerateKey(a.Size)
}

func (a *Aes) GenerateIV() (string, error) {
	return GenerateIV(16)
}

func (a *Aes) GenerateNonce() (string, error) {
	return GenerateNonce(12)
}

// reference: https://qa.1r1g.cn/crypto/ask/4985151/
func (a *Aes) GetNonce(key []byte) []byte {
	if a.IV != nil {
		return a.IV
	}

	return key[:12]
}

func New(size int, encoding Encoding, iv []byte) (*Aes, error) {
	if !isKeySizeAllow(size) {
		return nil, fmt.Errorf("aes key size should be 128, 192, 256, but %v", size)
	}

	if encoding == nil {
		encoding = &HexEncoding{}
	}

	if iv != nil && len(iv) != 16 {
		return nil, fmt.Errorf("aes iv size should be 128, but %v", len(iv)*8)
	}

	return &Aes{
		Size:     size,
		Encoding: encoding,
		IV:       iv,
	}, nil
}

func GenerateKey(size int) (string, error) {
	if !isKeySizeAllow(size) {
		return "", fmt.Errorf("aes secret size should be 128, 192, 256, but %v", size)
	}

	return RandomString(size / 8), nil
}

func GenerateIV(size int) (string, error) {
	if !isIVSizeAllow(size) {
		return "", fmt.Errorf("aes secret size should be 16, but %v", size)
	}

	return RandomString(size / 8), nil
}

func GenerateNonce(size int) (string, error) {
	if !isNonceSizeAllow(size) {
		return "", fmt.Errorf("aes secret size should be 12, but %v", size)
	}

	return RandomString(size / 8), nil
}
