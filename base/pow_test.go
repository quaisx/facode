package base

import (
	"testing"
)


func TestPow(t *testing.T) {
	if !VerifyHash("00123456789") {
		t.Error("Expected VerifyHash to return true")
	}
}