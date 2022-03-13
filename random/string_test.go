package random

import (
	"testing"
)

func Test_Random_String(t *testing.T) {
	length := 32
	randomStr := String(32)

	t.Log("random str:", randomStr)

	if len(randomStr) != length {
		t.Errorf("Expected length of %d, got %d", length, len(randomStr))
	}
}
