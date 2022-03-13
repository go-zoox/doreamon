package hmac

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func Sha256(secret, text string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(text))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func Md5(secret, text string) string {
	mac := hmac.New(md5.New, []byte(secret))
	mac.Write([]byte(text))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func Sha512(secret, text string) string {
	mac := hmac.New(sha512.New, []byte(secret))
	mac.Write([]byte(text))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func Sha1(secret, text string) string {
	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write([]byte(text))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func Sha224(secret, text string) string {
	mac := hmac.New(sha256.New224, []byte(secret))
	mac.Write([]byte(text))
	return fmt.Sprintf("%x", mac.Sum(nil))
}

func Sha384(secret, text string) string {
	mac := hmac.New(sha512.New384, []byte(secret))
	mac.Write([]byte(text))
	return fmt.Sprintf("%x", mac.Sum(nil))
}
