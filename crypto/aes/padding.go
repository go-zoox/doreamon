package aes

import "bytes"

type Padding interface {
	Padding(plainbytes []byte, blockSize int) []byte
	Unpadding(plainbytes []byte, blockSize int) []byte
}

type PCKS7Padding struct{}

func (p *PCKS7Padding) Padding(plainbytes []byte, blockSize int) []byte {
	paddingLen := blockSize - len(plainbytes)%blockSize
	padding := bytes.Repeat([]byte{byte(paddingLen)}, paddingLen)
	return append(plainbytes, padding...)
}

func (p *PCKS7Padding) Unpadding(plainbytes []byte, blockSize int) []byte {
	length := len(plainbytes)
	paddingLen := int(plainbytes[length-1])
	return plainbytes[:length-paddingLen]
}

type ZerosPadding struct{}

func (p *ZerosPadding) Padding(plainbytes []byte, blockSize int) []byte {
	paddingLen := blockSize - len(plainbytes)%blockSize
	padding := bytes.Repeat([]byte{byte(0)}, paddingLen)
	return append(plainbytes, padding...)
}

func (p *ZerosPadding) Unpadding(plainbytes []byte, blockSize int) []byte {
	// length := len(plainbytes)
	// paddingLen := 1
	// for index := length - 1; ; index -= 1 {
	// 	if int(plainbytes[length-1]) != 0 {
	// 		paddingLen = length - index
	// 		break
	// 	}
	// }

	// return plainbytes[:length-paddingLen]
	return bytes.TrimFunc(plainbytes, func(r rune) bool {
		return r == rune(0)
	})
}
