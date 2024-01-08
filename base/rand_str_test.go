package base

import (
	"testing"
)


func TestRndString(t *testing.T) {
	str := GenerateRandomString(10)
	if len(str) != 10 {
		t.Error("Generated string should be 10 characters")
	}
}