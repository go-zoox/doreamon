package crc64

import (
	"hash/crc64"
)

func Checksum(text string) uint64 {
	// c := crc64.New(crc64.MakeTable(crc64.ISO))
	// c.Write([]byte(text))
	// return c.Sum64()

	return crc64.Checksum([]byte(text), crc64.MakeTable(crc64.ISO))
}

func ChecksumSigned(text string) int64 {
	return int64(Checksum(text))
}
