package base

import (
    "crypto/sha256"
    "encoding/hex"
    "strconv"
)

const POW_DIFFICULTY = 2
const BYTE_ZERO = 48

// CalculateHash - generate a hash value for the challenge
// Returns:
// hash - SHA256 product for nonce + message
func CalculateHash(nonce int, message string) string {
    data := strconv.Itoa(nonce) + message
    hash := sha256.Sum256([]byte(data))
    return hex.EncodeToString(hash[:])
}

// VerifyHash - check if the hash has at least the correct number of leading zeros
func VerifyHash(hash string) bool {
	for _, c := range hash[:POW_DIFFICULTY] {
		if c != BYTE_ZERO {
			return false
		}
	}
	return true
}

// ProofOfWork - performs proof-of-work and solves the challenge
// Returns:
// - nonce that solves the challenge
func ProofOfWork(message string) (int, string) {
    nonce := 0
    for {
        hash := CalculateHash(nonce, message)
        if VerifyHash(hash) {
            return nonce, hash
        }
        nonce++
    }
}