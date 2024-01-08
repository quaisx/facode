# Test task for Server engineer (Go) - Design and implement “Word of Wisdom” tcp server

## 1. Problem

Design and implement “Word of Wisdom” tcp server:
+ TCP server should be protected from DDOS attacks with the Proof of Work (<https://en.wikipedia.org/wiki/Proof_of_work>), the challenge-response protocol should be used.
+ The choice of the POW algorithm should be explained.
+ After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.
+ Docker file should be provided both for the server and for the client that solves the POW challenge

## 2. Requirements

+ [Go 1.21+](https://go.dev/dl/) installed
+ [Docker](https://docs.docker.com/engine/install/) installed

## 3. Proof of Work

Proof of work (PoW) is a form of cryptographic proof in which one party (the prover) proves to others (the verifiers) that a certain amount of a specific computational effort has been expended. Verifiers can subsequently confirm this expenditure with minimal effort on their part. The concept was invented by Moni Naor and Cynthia Dwork in 1993 as a way to deter denial-of-service attacks and other service abuses such as spam on a network by requiring some work from a service requester, usually meaning processing time by a computer. The term "proof of work" was first coined and formalized in a 1999 paper by Markus Jakobsson and Ari Juels

### 3.1 Challenge-Response Protocol

Challenge–response protocols assume a direct interactive link between the requester (client) and the provider (server). The provider chooses a challenge, say an item in a set with a property, the requester finds the relevant response in the set, which is sent back and checked by the provider. As the challenge is chosen on the spot by the provider, its difficulty can be adapted to its current load. The work on the requester side may be bounded if the challenge-response protocol has a known solution (chosen by the provider), or is known to exist within a bounded search space.

### 3.2 PoW algorithms

Integer square root modulo a large prime
Weaken Fiat–Shamir signatures
Ong–Schnorr–Shamir signature broken by Pollard
Partial hash inversion
Hash sequences
Puzzles
Diffie-Hellman–based puzzle
Moderate
Mbound
Hokkaido
Cuckoo Cycle
Merkle tree–based
Guided tour puzzle protocol

### 3.3 Algorithm choice

The algorithm of choice for this assignment has been contrived in-house to simplify the pow work as well as to demonstrate the line of thinking when dealing with pow. The algorithm is based on the SHA-256 hashing function which calculates the hash product from the compounded input Nonce+challenge. When the server receives a new request from a client, it generates a random challenge phrase, caches it and sends it to the client. The client has to find a nonce value that, when combined with the challenge phrase, produces the desired number of leading zeros. For the sake of simplicity the difficulty level has been set to TWO leading zeros.

## 4. How to test

To run the tests (for brevity, the tests are oversimplified):

```bash
make test
```

To run the server:

```bash
make run-server
```

To run the client:

```bash
make run-client
```

## 5. What to evolve next

+ The protocol logic should have a state machine which determines the transition rules between states
+ Caching logic should have expiration time which limits the client from taking indefinitely to complete the challenge
+ Retrying logic should be implemented
+ CLI should be added to both the client and server to change the default behaviours
+ Better test coverage; The current implementation only acknowledges that the tests are required
+ Better logic to blacklist ill behaved clients
