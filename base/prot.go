package base

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	PROT_CS_Proof = iota
	PROT_SC_Proof
	PROT_CS_Quote
	PROT_SC_Quote
)

// message field seprator
const PROT_MSG_SEP = "|"

type ProtMessage struct {
	Type  int    
	Payload string 
}

func (m *ProtMessage) String() string {
	return fmt.Sprintf("%d|%s", m.Type, m.Payload)
}

// ProtParseMessage - parse the client/server exchange message
func ProtParseMessage(str string) (*ProtMessage, error) {
	str = strings.TrimSpace(str)
	var msgType int
	// split the message into parts: type and payload
	parts := strings.Split(str, PROT_MSG_SEP)
	if len(parts) < 1 || len(parts) > 2 {
		return nil, fmt.Errorf("message does not comply with the protocol")
	}
	// obtain message type first
	msgType, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("cannot parse header message type")
	}
	msg := &ProtMessage{
		Type: msgType,
	}
	// the part after | is payload
	if len(parts) == 2 {
		msg.Payload = parts[1]
	}
	return msg, nil
}
