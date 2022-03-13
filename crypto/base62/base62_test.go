package base62

import (
	"testing"

	"github.com/go-zoox/doreamon/test"
)

func TestBase64(t *testing.T) {
	ts := test.TestSuit{T: t}

	pairs := []struct {
		source int
		target string
	}{
		{
			source: 65535,
			target: "h31",
		},
		{
			source: 12345,
			target: "3d7",
		},
		{
			source: 123456,
			target: "w7e",
		},
		{
			source: 1234567,
			target: "5ban",
		},
		{
			source: 12345678,
			target: "PNFQ",
		},
		{
			source: 123456789,
			target: "8m0Kx",
		},
		{
			source: 1234567890,
			target: "1ly7vk",
		},
		{
			source: 12345678901,
			target: "dtvd3f",
		},
	}

	for _, pair := range pairs {
		ts.Expect(Encode(pair.source)).ToEqual(pair.target)
		ts.Expect(Decode(pair.target)).ToEqual(pair.source)
	}
}
