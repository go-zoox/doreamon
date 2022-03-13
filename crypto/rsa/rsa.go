package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"errors"
)

type RSA struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func New(privateKey, publicKey string) (*RSA, error) {
	privateKeyDecoded, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return nil, errors.New("invalid private key(1): " + err.Error())
	}

	_privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyDecoded)
	if err != nil {
		return nil, errors.New("invalid private key(2): " + err.Error())
	}

	publicKeyDecoded, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, errors.New("invalid public key(1): " + err.Error())
	}

	_publicKey, err := x509.ParsePKCS1PublicKey(publicKeyDecoded)
	if err != nil {
		return nil, errors.New("invalid public key(2): " + err.Error())
	}

	return &RSA{
		privateKey: _privateKey,
		publicKey:  _publicKey,
	}, nil
}

func (r *RSA) Encrypt(plainbytes []byte) (cipherbytes []byte, err error) {
	cipherbytes, err = rsa.EncryptPKCS1v15(rand.Reader, r.publicKey, plainbytes)
	return
}

func (r *RSA) Decrypt(cipherbytes []byte) (plainbytes []byte, err error) {
	plainbytes, err = rsa.DecryptPKCS1v15(rand.Reader, r.privateKey, cipherbytes)
	return
}

func (r *RSA) Sign(message []byte) (signature []byte, err error) {
	hash := sha256.New()
	_, err = hash.Write(message)
	if err != nil {
		panic(err)
	}
	sum := hash.Sum(nil)

	signature, err = rsa.SignPKCS1v15(rand.Reader, r.privateKey, crypto.SHA256, sum)
	return
}

func (r *RSA) Verify(message, signature []byte) (ok bool, err error) {
	hash := sha256.New()
	_, err = hash.Write(message)
	if err != nil {
		panic(err)
	}
	sum := hash.Sum(nil)

	err = rsa.VerifyPKCS1v15(r.publicKey, crypto.SHA256, sum, signature)
	if err != nil {
		return
	}

	return true, nil
}
