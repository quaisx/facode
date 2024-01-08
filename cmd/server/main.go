// server.go
package main

import (
	"bufio"
	"facode/base"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

var quotes = []string{
	"The only true wisdom is in knowing you know nothing. - Socrates",
	"In three words I can sum up everything I've learned about life: it goes on. - Robert Frost",
	"I know that I am intelligent because I know that I know nothing. - Socrates",
	"To find yourself, think for yourself. - Socrates",
	"Courage is knowing what not to fear. - Plato",
}

const CHALLENGE_LENGTH = 24 // string length for the challenge string
var cache = sync.Map{}      // client connection cache
var blacklist = sync.Map{}  // blacklist ill behaved clients

// handleConnection - communication steps between the server and the client
func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		//server receives a request from client
		req, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("client %s closed the connection: %v\n", conn.RemoteAddr().String(), err)
			return
		}
		// parse client request, generate a response message
		msg, err := processRequest(req, conn.RemoteAddr().String())
		if err != nil {
			fmt.Printf("error processing request: %v", err)
			return
		}
		// send the response back to the server
		if msg != nil {
			err := sendMsg(*msg, conn)
			if err != nil {
				fmt.Println("error sending response message:", err)
			}
		}
	}
}

func main() {
	// start listening for new connections
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("error starting server: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Println("server started. Listening on :8080")

	for {
		// accept new client connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("error accepting connection: %v", err)
			continue
		}
		// check if the client has previously been blacklisted
		_, ok := blacklist.Load(conn.RemoteAddr().String())
		if ok {
			fmt.Printf("Ignore a connection from a blacklisted client %s.\n", conn.RemoteAddr().String())
			conn.Close()
			continue
		}
		// handle each connection in a separate goroutine
		go handleConnection(conn)
	}
}

// processRequest - handle client requests
func processRequest(msgStr string, clientAddr string) (*base.ProtMessage, error) {
	msg, err := base.ProtParseMessage(msgStr)
	if err != nil {
		return nil, err
	}
	// switch by type of request msg
	switch msg.Type {
	// Client to server request to generate a new challenge
	case base.PROT_CS_Proof:
		fmt.Printf("client %s requests pow challenge\n", clientAddr)
		// create new challenge for client
		challenge := base.GenerateRandomString(CHALLENGE_LENGTH)
		cache.Store(clientAddr, challenge)
		msg := &base.ProtMessage{
			Type:    base.PROT_SC_Proof,
			Payload: challenge,
		}
		return msg, nil
	// Client to server request to generate a quote given a valid pow
	case base.PROT_CS_Quote:
		fmt.Printf("client %s requests resource with nonce %s as solution. Verifying...\n", clientAddr, msg.Payload)
		// local cache lookup for the generated challenge
		client_challenge, ok := cache.Load(clientAddr)
		if !ok {
			// this client is not known to the server, blacklist it
			blacklist.Store(clientAddr, time.Now().UnixMilli())
			return nil, fmt.Errorf("error client %s protocol violation", clientAddr)
		}
		// client found a nonce that solves the challenge
		nonce, err := strconv.Atoi(msg.Payload)
		if err != nil {
			return nil, fmt.Errorf("error obtaining nonce: %v", err)
		}
		challenge := client_challenge.(string)
		// check of valid pow solution
		solved_challenge := base.CalculateHash(nonce, challenge)
		if !base.VerifyHash(solved_challenge) {
			// if client fails to verify with the server, it has to obtain a new challenge
			cache.Delete(clientAddr)
			return nil, fmt.Errorf("error client %s failed to solve the challenge", clientAddr)
		}

		//get random quote
		fmt.Printf("client %s successfully computed pow for %s+%s\n", clientAddr, msg.Payload, challenge)
		msg := &base.ProtMessage{
			Type:    base.PROT_SC_Quote,
			Payload: quotes[nonce%len(quotes)],
		}
		// remove the client from the cache after a quote has been served
		cache.Delete(clientAddr)
		return msg, nil
	default:
		return nil, fmt.Errorf("unknown request type")
	}
}

// sendMsg - send a message to the client
func sendMsg(msg base.ProtMessage, conn net.Conn) error {
	msgStr := fmt.Sprintf("%s\n", msg.String())
	_, err := conn.Write([]byte(msgStr))
	return err
}
