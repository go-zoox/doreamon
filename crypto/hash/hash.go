package hash

import (
	_md5 "crypto/md5"
	_sha1 "crypto/sha1"
	_sha256 "crypto/sha256"
	_sha512 "crypto/sha512"
	"fmt"
	"io"

	"github.com/go-zoox/doreamon/crypto/base62"
	"github.com/spaolacci/murmur3"
)

func Md5(text string) string {
	h := _md5.New()
	io.WriteString(h, text)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Sha256(text string) string {
	h := _sha256.New()
	io.WriteString(h, text)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Sha512(text string) string {
	h := _sha512.New()
	io.WriteString(h, text)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Sha1(text string) string {
	h := _sha1.New()
	io.WriteString(h, text)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Sha224(text string) string {
	h := _sha256.New224()
	io.WriteString(h, text)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Sha384(text string) string {
	h := _sha512.New384()
	io.WriteString(h, text)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func MurmurHash(text string) string {
	m := murmur3.Sum32([]byte(text))
	return base62.Encode(int(m))
}
