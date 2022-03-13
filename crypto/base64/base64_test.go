package base64

import (
	"testing"

	"github.com/go-zoox/doreamon/test"
)

func TestBase64(t *testing.T) {
	ts := test.TestSuit{T: t}

	source := "helloworld"
	target := "aGVsbG93b3JsZA=="

	ts.Expect(Encode(source)).ToEqual(target)
	ts.Expect(Decode(target)).ToEqual(source)
}
