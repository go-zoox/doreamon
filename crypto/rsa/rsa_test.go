package rsa

import (
	"encoding/base64"
	"testing"
)

func TestRsaGeneratePrivateKey(t *testing.T) {
	privateKey, err := GeneratePrivateKeyString(1024)
	if err != nil {
		t.Fatal("error:", err)
	}

	if privateKey == "" {
		t.Fatal("failed to generate private key")
	}
}

func TestRsaGeneratePublicKey(t *testing.T) {
	privateKey := "MIICXAIBAAKBgQCmuEbvwpNB//67t/2g5cGMFkNbkmRbRmtDBK+hjboF66ml7hdbYPF09GNHQbl7b8Ru1hWTNhtu37GF4X0zg8nIU0HSMRSfvzUJ2SWKoAKXPy0jYQk2pxYpkYr3RfMMYVaEa55sT+0MGipSmpibrwkUL6W7k7CTYZpd/9J3JjAWRwIDAQABAoGANo0tiN4d2QaujzXQ44jKH9BZEemAtO0Bw9gQr8f0CmPmCskxE1FRMHeW1IYI7v7PQ4UBYj3eFBVVvPzfPq/sofxQnwVYVSVjWgz6NdZsaTAFh1YxnJx/IzAQFeWFyigZbmSBrMtLAer2G6inEOornzqT0+n8GEkeOpG+h7s54iECQQDaxwtpa11fiZov7dKdFJeOsYoGizBHafkA3/PrnRZxjhj+orWy87ev8Ltp+rz/5JnrHn7Pq31hgDn8LIWRiCyLAkEAwxXWCYTwE3N6KZ7UgxHpxBOaKyiQBLUfhu9rMeSyM4xdgbT6ByEwPjJxuBrqFQoaBUSLoX6vFGohJkFoUdTItQJACgctortlIEfyZVgFW2XiPIwuw3YF1IArBbs+NwKQUMwuoR1cLsO1G79xF76Cg0g7NefD8EjwClQSVFjGFpGjWQJAXcE4xApndnGg3C/A4dzSA7GH/gXYcOq65BZb5faKzcs/hP58ysBgdwO3M0t8A/B+4Nk4YbyIV79JfyEgCXPBoQJBALMJZOROVZZND4dUQAxk2+aRR+JIC7R1VDvNiUsQTiv9BRIs7l4qAiwCuDpdIL7y9t2AO8kc+5wINkjJUs4dq1Q="
	publickKey := "MIGJAoGBAKa4Ru/Ck0H//ru3/aDlwYwWQ1uSZFtGa0MEr6GNugXrqaXuF1tg8XT0Y0dBuXtvxG7WFZM2G27fsYXhfTODychTQdIxFJ+/NQnZJYqgApc/LSNhCTanFimRivdF8wxhVoRrnmxP7QwaKlKamJuvCRQvpbuTsJNhml3/0ncmMBZHAgMBAAE="

	_publicKey, err := GeneratePublicKeyString(privateKey)
	if err != nil {
		t.Fatal("error:", err)
	}

	if _publicKey != publickKey {
		t.Fatal("failed to generate public key from private key")
	}

	// fmt.Println(publickKey)
	// fmt.Println("xxx")
	// fmt.Println(_publicKey)

	// ts := test.TestSuit{}
	// ts.Expect(_publicKey).ToEqual(publickKey)
}

func TestEncryptDecrypt(t *testing.T) {
	privateKey := "MIICXAIBAAKBgQCmuEbvwpNB//67t/2g5cGMFkNbkmRbRmtDBK+hjboF66ml7hdbYPF09GNHQbl7b8Ru1hWTNhtu37GF4X0zg8nIU0HSMRSfvzUJ2SWKoAKXPy0jYQk2pxYpkYr3RfMMYVaEa55sT+0MGipSmpibrwkUL6W7k7CTYZpd/9J3JjAWRwIDAQABAoGANo0tiN4d2QaujzXQ44jKH9BZEemAtO0Bw9gQr8f0CmPmCskxE1FRMHeW1IYI7v7PQ4UBYj3eFBVVvPzfPq/sofxQnwVYVSVjWgz6NdZsaTAFh1YxnJx/IzAQFeWFyigZbmSBrMtLAer2G6inEOornzqT0+n8GEkeOpG+h7s54iECQQDaxwtpa11fiZov7dKdFJeOsYoGizBHafkA3/PrnRZxjhj+orWy87ev8Ltp+rz/5JnrHn7Pq31hgDn8LIWRiCyLAkEAwxXWCYTwE3N6KZ7UgxHpxBOaKyiQBLUfhu9rMeSyM4xdgbT6ByEwPjJxuBrqFQoaBUSLoX6vFGohJkFoUdTItQJACgctortlIEfyZVgFW2XiPIwuw3YF1IArBbs+NwKQUMwuoR1cLsO1G79xF76Cg0g7NefD8EjwClQSVFjGFpGjWQJAXcE4xApndnGg3C/A4dzSA7GH/gXYcOq65BZb5faKzcs/hP58ysBgdwO3M0t8A/B+4Nk4YbyIV79JfyEgCXPBoQJBALMJZOROVZZND4dUQAxk2+aRR+JIC7R1VDvNiUsQTiv9BRIs7l4qAiwCuDpdIL7y9t2AO8kc+5wINkjJUs4dq1Q="
	publickKey := "MIGJAoGBAKa4Ru/Ck0H//ru3/aDlwYwWQ1uSZFtGa0MEr6GNugXrqaXuF1tg8XT0Y0dBuXtvxG7WFZM2G27fsYXhfTODychTQdIxFJ+/NQnZJYqgApc/LSNhCTanFimRivdF8wxhVoRrnmxP7QwaKlKamJuvCRQvpbuTsJNhml3/0ncmMBZHAgMBAAE="

	source := "helloworld"
	// target := "ZUiH2GD7Jp1XKIvTxurEelJiv7Ke1/ZacZGiMKXrU4tyCfhz5OeiYho2TPrH/dEZ25JjoRV9CGU38/8AP+l+/eeqGHlezVbGlfUM4yK4S9ELPr9OpGqT7H0+tpwZX+5HsU1eDdzXPL+UQpMQ4joRM9c1lgO+z4MmMC0sPHtWdy4="

	r, err := New(privateKey, publickKey)
	if err != nil {
		t.Fatal("error:", err)
	}

	// if _target, err := r.Encrypt([]byte(source)); err != nil || base64.StdEncoding.EncodeToString(_target) != target {
	// 	t.Fatal("failed to encrypt: " + base64.StdEncoding.EncodeToString(_target))
	// }

	target, err := r.Encrypt([]byte(source))
	if err != nil {
		t.Fatal("failed to encrypt: " + err.Error())
	}

	_source, err := r.Decrypt([]byte(target))
	if err != nil {
		t.Fatal("failed to decrypt: " + err.Error())
	}

	if source != string(_source) {
		t.Fatal("failed to rsa")
	}
}

func TestSignature(t *testing.T) {
	privateKey := "MIICXAIBAAKBgQCmuEbvwpNB//67t/2g5cGMFkNbkmRbRmtDBK+hjboF66ml7hdbYPF09GNHQbl7b8Ru1hWTNhtu37GF4X0zg8nIU0HSMRSfvzUJ2SWKoAKXPy0jYQk2pxYpkYr3RfMMYVaEa55sT+0MGipSmpibrwkUL6W7k7CTYZpd/9J3JjAWRwIDAQABAoGANo0tiN4d2QaujzXQ44jKH9BZEemAtO0Bw9gQr8f0CmPmCskxE1FRMHeW1IYI7v7PQ4UBYj3eFBVVvPzfPq/sofxQnwVYVSVjWgz6NdZsaTAFh1YxnJx/IzAQFeWFyigZbmSBrMtLAer2G6inEOornzqT0+n8GEkeOpG+h7s54iECQQDaxwtpa11fiZov7dKdFJeOsYoGizBHafkA3/PrnRZxjhj+orWy87ev8Ltp+rz/5JnrHn7Pq31hgDn8LIWRiCyLAkEAwxXWCYTwE3N6KZ7UgxHpxBOaKyiQBLUfhu9rMeSyM4xdgbT6ByEwPjJxuBrqFQoaBUSLoX6vFGohJkFoUdTItQJACgctortlIEfyZVgFW2XiPIwuw3YF1IArBbs+NwKQUMwuoR1cLsO1G79xF76Cg0g7NefD8EjwClQSVFjGFpGjWQJAXcE4xApndnGg3C/A4dzSA7GH/gXYcOq65BZb5faKzcs/hP58ysBgdwO3M0t8A/B+4Nk4YbyIV79JfyEgCXPBoQJBALMJZOROVZZND4dUQAxk2+aRR+JIC7R1VDvNiUsQTiv9BRIs7l4qAiwCuDpdIL7y9t2AO8kc+5wINkjJUs4dq1Q="
	publickKey := "MIGJAoGBAKa4Ru/Ck0H//ru3/aDlwYwWQ1uSZFtGa0MEr6GNugXrqaXuF1tg8XT0Y0dBuXtvxG7WFZM2G27fsYXhfTODychTQdIxFJ+/NQnZJYqgApc/LSNhCTanFimRivdF8wxhVoRrnmxP7QwaKlKamJuvCRQvpbuTsJNhml3/0ncmMBZHAgMBAAE="

	message := "helloworld"
	signature := "FI3pnHIadq5zK7mrmLf/uvW/i2GzBg18O2wA1QZqgGKaAdKLrllSnlSOVBeu8GqgoQCPZWdWzFPlCl40GbgRrqLt3qsp0gw6z5EWNcQ1stXfWATVReBBXClxT0FoF7jGk6kKqf118SGc1/b1PHokoXWDI9V5iETmZofd6VSPWGA="

	r, err := New(privateKey, publickKey)
	if err != nil {
		t.Fatal("error:", err)
	}

	_signature, err := r.Sign([]byte(message))
	if err != nil {
		t.Fatal("error:", err)
	}

	if base64.StdEncoding.EncodeToString([]byte(_signature)) != signature {
		t.Fatal("failed to sign")
	}

	if ok, err := r.Verify([]byte(message), _signature); !ok || err != nil {
		t.Fatal("failed to verify")
	}
}
