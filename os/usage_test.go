package os

import (
	"testing"
)

func Test_OS_Usage(t *testing.T) {
	usage := GetOSUsage()
	if usage == nil {
		t.Error("failed to test usage")
	}

	t.Log("usage:", usage)
}
