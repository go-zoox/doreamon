package shorturl

import "github.com/go-zoox/crypto/hash"

func ShortURL(longUrl string) string {
	return hash.MurmurHash(longUrl)
}
