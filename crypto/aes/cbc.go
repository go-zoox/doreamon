// reference:
//	https://blog.csdn.net/wade3015/article/details/84454836
package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

type CBC struct {
	Aes
	Padding Padding
}

func NewCBC(size int, padding Padding, encoding Encoding, iv []byte) (*CBC, error) {
	aes, err := New(size, encoding, iv)
	if err != nil {
		return nil, err
	}

	if padding == nil {
		padding = &PCKS7Padding{}
	}

	return &CBC{
		Aes:     *aes,
		Padding: padding,
	}, nil
}

func (a *CBC) Encrypt(plainbytes, key []byte) (cipherbytes []byte, err error) {
	secretLength := len(key)
	if secretLength != a.BlockSize() {
		err = fmt.Errorf("secret size should be %v, but %v", a.BlockSize(), secretLength)
		return
	}

	var block cipher.Block
	block, err = aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCEncrypter(block, a.GetIV(key))

	_plainbytes := a.Padding.Padding(plainbytes, block.BlockSize())
	_cipherbytes := make([]byte, len(_plainbytes))

	blockMode.CryptBlocks(_cipherbytes, _plainbytes)
	// encoding
	cipherbytes = []byte(a.Encoding.EncodeToString(_cipherbytes))
	return
}

func (a *CBC) Decrypt(cipherbytes, key []byte) (plainbytes []byte, err error) {
	secretLength := len(key)
	if secretLength != a.BlockSize() {
		err = fmt.Errorf("secret size should be %v, but %v", a.BlockSize(), secretLength)
		return
	}

	var block cipher.Block
	block, err = aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, a.GetIV(key))

	_cipherbytes, err := a.Encoding.DecodeString(string(cipherbytes))
	if err != nil {
		return nil, err
	}

	_plainbytes := make([]byte, len(_cipherbytes))
	blockMode.CryptBlocks(_plainbytes, _cipherbytes)
	plainbytes = a.Padding.Unpadding(_plainbytes, block.BlockSize())
	return
}

func (a *CBC) EncryptString(plaintext, secret string) (ciphertext string, err error) {
	cipherbytes, err := a.Encrypt([]byte(plaintext), []byte(secret))
	if err != nil {
		return "", err
	}

	return string(cipherbytes), nil
}

func (a *CBC) DecryptString(ciphertext, secret string) (plaintext string, err error) {
	plainbytes, err := a.Decrypt([]byte(ciphertext), []byte(secret))
	if err != nil {
		return "", err
	}

	return string(plainbytes), nil
}
