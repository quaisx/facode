package base

import (
  "math/rand"
  "time"
)

// charset to use when generating random strings
const charset = "abcdefghijklmnopqrstuvwxyz" +
  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

  var seededRand *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))

// StringWithCharset returns a randomly generated string
// Parameters:
//  length - the desired length of the string
//  charset - the charset to use when generating random strings
  func StringWithCharset(length int, charset string) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

func GenerateRandomString(length int) string {
  return StringWithCharset(length, charset)
}