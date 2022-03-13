package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

func GeneratePrivateKey(bits int) ([]byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	return derStream, nil
}

func GeneratePublicKey(privateKey []byte) ([]byte, error) {
	_privateKey, err := x509.ParsePKCS1PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	publicKey := &_privateKey.PublicKey
	return x509.MarshalPKCS1PublicKey(publicKey), nil
}

func GeneratePrivateKeyString(bits int) (string, error) {
	stream, err := GeneratePrivateKey(bits)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(stream), nil
}

func GeneratePublicKeyString(privateKey string) (string, error) {
	_privateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", err
	}

	stream, err := GeneratePublicKey(_privateKey)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(stream), nil
}

func GeneratePrivateKeyFile(bits int, filepath string) error {
	stream, err := GeneratePrivateKey(bits)
	if err != nil {
		return err
	}

	return PEMEncode(filepath, stream, true)
}

func GeneratePublicKeyFile(privateKey []byte, filepath string) error {
	stream, err := GeneratePublicKey(privateKey)
	if err != nil {
		return err
	}

	return PEMEncode(filepath, stream, false)
}
