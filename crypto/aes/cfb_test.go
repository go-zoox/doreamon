package aes

import (
	"testing"

	"github.com/go-zoox/doreamon/test"
)

func Test_AES_CFB(t *testing.T) {
	testdata := "helloworld"
	testcases := map[string]string{
		// aes-128-cfb
		"mysecretmysecret": "B1SLOXpNNrKSdQ==",
		// aes-192-cfb
		"mysecretmysecretmysecret": "OBGDXZigSAWEgA==",
		// aes-256-cfb
		"mysecretmysecretmysecretmysecret": "F/uw+MDTPx5EPg==",
	}

	for secret, encrypted := range testcases {
		aes, err := NewCFB(len(secret)*8, &Base64Encoding{}, nil)
		if err != nil {
			t.Fatalf("AES New failed: %s", err.Error())
		}

		ts := test.TestSuit{T: t}

		enc, err := aes.EncryptString(testdata, secret)
		if err != nil {
			t.Fatalf("AES Encrypt failed: %s", err.Error())
		}
		dec, _ := aes.DecryptString(enc, secret)

		// ts.Expect(base64.StdEncoding.EncodeToString(enc)).ToEqual(data.expected)
		ts.Expect(enc).ToEqual(encrypted)
		ts.Expect(dec).ToEqual(testdata)
	}
}
