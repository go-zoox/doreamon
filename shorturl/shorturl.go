package shorturl

import "github.com/go-zoox/doreamon/crypto/hash"

func ShortURL(longUrl string) string {
	return hash.MurmurHash(longUrl)
}
