// reference:
//	https://gist.github.com/temoto/5052503
package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

type CFB struct {
	Aes
}

func NewCFB(size int, encoding Encoding, iv []byte) (*CFB, error) {
	aes, err := New(size, encoding, iv)
	if err != nil {
		return nil, err
	}

	return &CFB{
		Aes: *aes,
	}, nil
}

func (a *CFB) Encrypt(plainbytes, key []byte) (cipherbytes []byte, err error) {
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

	blockMode := cipher.NewCFBEncrypter(block, a.GetIV(key))
	_cipherbytes := make([]byte, len(plainbytes))

	blockMode.XORKeyStream(_cipherbytes, plainbytes)
	// encoding
	cipherbytes = []byte(a.Encoding.EncodeToString(_cipherbytes))
	return
}

func (a *CFB) Decrypt(cipherbytes, key []byte) (plainbytes []byte, err error) {
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
	blockMode := cipher.NewCFBDecrypter(block, a.GetIV(key))

	_cipherbytes, err := a.Encoding.DecodeString(string(cipherbytes))
	if err != nil {
		return nil, err
	}

	plainbytes = make([]byte, len(_cipherbytes))
	blockMode.XORKeyStream(plainbytes, _cipherbytes)
	return
}

func (a *CFB) EncryptString(plaintext, secret string) (ciphertext string, err error) {
	cipherbytes, err := a.Encrypt([]byte(plaintext), []byte(secret))
	if err != nil {
		return "", err
	}

	return string(cipherbytes), nil
}

func (a *CFB) DecryptString(ciphertext, secret string) (plaintext string, err error) {
	plainbytes, err := a.Decrypt([]byte(ciphertext), []byte(secret))
	if err != nil {
		return "", err
	}

	return string(plainbytes), nil
}
