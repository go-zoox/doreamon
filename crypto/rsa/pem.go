package rsa

import (
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func PEMEncode(filepath string, stream []byte, isPrivate bool) error {
	typ := "PRIVATE"
	if !isPrivate {
		typ = "PUBLIC"
	}

	block := &pem.Block{
		Type:  fmt.Sprintf("RSA %s KEY", typ),
		Bytes: stream,
	}

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = pem.Encode(file, block)
	if err != nil {
		return err
	}

	return nil
}

func PEMDecode(filepath string) (*pem.Block, error) {
	fs, err := os.Open(filepath)
	if err != nil {
		return nil, errors.New("failed to load key: " + err.Error())
	}

	key, err := ioutil.ReadAll(fs)
	if err != nil {
		return nil, errors.New("failed to read key: " + err.Error())
	}

	block, _ := pem.Decode(key)
	if block == nil {
		return nil, errors.New("failed to decode key")
	}

	return block, nil
}
