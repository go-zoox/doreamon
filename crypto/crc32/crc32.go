package crc32

import (
	"hash/crc32"
)

func Checksum(text string) uint32 {
	// c := crc32.NewIEEE()
	// c.Write([]byte(text))
	// return c.Sum32()

	return crc32.ChecksumIEEE([]byte(text))

	// c := crc32.New(crc32.MakeTable(crc32.Koopman))
	// c.Write([]byte(text))
	// return c.Sum32()
}

func ChecksumSigned(text string) int32 {
	return int32(Checksum(text))
}
