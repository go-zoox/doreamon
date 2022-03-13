package md5

import (
	_md5 "crypto/md5"
	"fmt"
	"io"
)

func Md5(text string) string {
	h := _md5.New()
	io.WriteString(h, text)
	return fmt.Sprintf("%x", h.Sum(nil))
}
