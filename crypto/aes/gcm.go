// reference:
//	https://github.com/SimonWaldherr/golang-examples/blob/master/advanced/aesgcm.go
package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

type GCM struct {
	Aes
}

func NewGCM(size int, encoding Encoding, iv []byte) (*GCM, error) {
	aes, err := New(size, encoding, iv)
	if err != nil {
		return nil, err
	}

	return &GCM{
		Aes: *aes,
	}, nil
}

func (a *GCM) Encrypt(plainbytes, key []byte) (cipherbytes []byte, err error) {
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

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	_cipherbytes := gcm.Seal(nil, a.GetNonce(key), plainbytes, nil)

	// encoding
	cipherbytes = []byte(a.Encoding.EncodeToString(_cipherbytes))
	return
}

func (a *GCM) Decrypt(cipherbytes, key []byte) (plainbytes []byte, err error) {
	_cipherbytes, _ := a.Encoding.DecodeString(string(cipherbytes))

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

	var gcm cipher.AEAD
	gcm, err = cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plainbytes, err = gcm.Open(nil, a.GetNonce(key), _cipherbytes, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (a *GCM) EncryptString(plaintext, secret string) (ciphertext string, err error) {
	cipherbytes, err := a.Encrypt([]byte(plaintext), []byte(secret))
	if err != nil {
		return "", err
	}

	return string(cipherbytes), nil
}

func (a *GCM) DecryptString(ciphertext, secret string) (plaintext string, err error) {
	plainbytes, err := a.Decrypt([]byte(ciphertext), []byte(secret))
	if err != nil {
		return "", err
	}

	return string(plainbytes), nil
}
