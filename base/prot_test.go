package base

import (
	"testing"
)


func TestProt(t *testing.T) {
	m, e := ProtParseMessage("0|PROT_CS_Proof")
	if e != nil {
		t.Errorf("Not expected error: %v", e)
	}
	if m.Type != PROT_CS_Proof {
		t.Errorf("Expected %v, got %v", PROT_CS_Proof, m.Type)
	}
}