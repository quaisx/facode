// client.go
package main

import (
	"bufio"
	"facode/base"
	"fmt"
	"net"
	"strconv"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	msgProt, err := processConnection(conn)
	if err != nil {
		fmt.Printf("error processing server connection: %v\n", err)
		return
	}
	if msgProt != nil {
		fmt.Printf("Server says: %s\n", msgProt.Payload)
	}
}

// processConnection - communication steps between the client and the server
func processConnection(conn net.Conn) (*base.ProtMessage, error) {

	reader := bufio.NewReader(conn)

	req_challenge := &base.ProtMessage{
		Type: base.PROT_CS_Proof,
	}
	// request a challenge
	err := sendMsg(req_challenge, conn)
	if err != nil {		
		return nil, fmt.Errorf("error sending challenge: %v", err)
	}
	// receive a new challenge to solve
	msgRcv, err := rcvMsg(reader)
	if err != nil {
		return nil, fmt.Errorf("error reading server pow response: %v", err)
	}
	// parsing the challenge response
	msgProt, err := base.ProtParseMessage(msgRcv)
	if err != nil {	
		return nil, fmt.Errorf("error parsing server response message: %v", err)
	}

	// perform Proof of Work
	nonce, _ := base.ProofOfWork(msgProt.Payload)
	// send a request for quote along with the nonce value that solves the given pow
	msg := &base.ProtMessage{
		Type:    base.PROT_CS_Quote,
		Payload: strconv.Itoa(nonce),
	}
	// send the quote request to the server
	err = sendMsg(msg, conn)
	if err != nil {		
		return nil, fmt.Errorf("error sending quote request: %v", err)
	}
	// receive a quote response from the server
	msgRcv, err = rcvMsg(reader)
	if err != nil {		
		return nil, fmt.Errorf("error reading server quote response: %v", err)
	}

	// parsing the quote response
	msgProt, err = base.ProtParseMessage(msgRcv)
	if err != nil {	
		return nil, fmt.Errorf("error parsing server response message: %v", err)
	}	
	return msgProt, nil
}

// rcvMsg - client receive message
func rcvMsg(r *bufio.Reader) (string, error) {
	return r.ReadString('\n')
}

// sendMsg - client send message
func sendMsg(msg *base.ProtMessage, conn net.Conn) error {
	msgStr := fmt.Sprintf("%s\n", msg)
	_, err := conn.Write([]byte(msgStr))
	return err
}
