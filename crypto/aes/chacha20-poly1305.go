// reference:
//	https://github.com/SimonWaldherr/golang-examples/blob/master/advanced/aesChacha20Poly1305.go
package aes

import (
	"fmt"

	chacha20poly1305 "golang.org/x/crypto/chacha20poly1305"
)

type Chacha20Poly1305 struct {
	Aes
}

func NewChacha20Poly1305(encoding Encoding, iv []byte) (*Chacha20Poly1305, error) {
	aes, err := New(chacha20poly1305.KeySize*8, encoding, iv)
	if err != nil {
		return nil, err
	}

	return &Chacha20Poly1305{
		Aes: *aes,
	}, nil
}

func (a *Chacha20Poly1305) Encrypt(plainbytes, key []byte) (cipherbytes []byte, err error) {
	if len(key) != chacha20poly1305.KeySize {
		err = fmt.Errorf("secret size should be %v, but %v", chacha20poly1305.KeySize, len(key))
		return
	}

	ahead, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, err
	}

	nonce := a.GetNonce(key)
	_cipherbytes := ahead.Seal(nil, nonce, plainbytes, nil)

	// encoding
	cipherbytes = []byte(a.Encoding.EncodeToString(_cipherbytes))
	return
}

func (a *Chacha20Poly1305) Decrypt(cipherbytes, key []byte) (plainbytes []byte, err error) {
	_cipherbytes, _ := a.Encoding.DecodeString(string(cipherbytes))

	if len(key) != chacha20poly1305.KeySize {
		err = fmt.Errorf("secret size should be %v, but %v", chacha20poly1305.KeySize, len(key))
		return
	}

	ahead, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, err
	}

	nonce := a.GetNonce(key)
	plainbytes, err = ahead.Open(nil, nonce, _cipherbytes, nil)
	if err != nil {
		return nil, err
	}

	return
}

func (a *Chacha20Poly1305) EncryptString(plaintext, secret string) (ciphertext string, err error) {
	cipherbytes, err := a.Encrypt([]byte(plaintext), []byte(secret))
	if err != nil {
		return "", err
	}

	return string(cipherbytes), nil
}

func (a *Chacha20Poly1305) DecryptString(ciphertext, secret string) (plaintext string, err error) {
	plainbytes, err := a.Decrypt([]byte(ciphertext), []byte(secret))
	if err != nil {
		return "", err
	}

	return string(plainbytes), nil
}
