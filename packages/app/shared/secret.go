package shared

import (
	"crypto/rand"
	"io"
	"log"
	"os"
)

func init() {
	assertAvailablePRNG()
}

func assertAvailablePRNG() {
	// Assert that a cryptographically secure PRNG is available.
	// Panic otherwise.
	buf := make([]byte, 1)

	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		log.Fatalf("crypto/rand is unavailable: Read() failed with %#v", err)
	}
}

// GenerateRandomBytes returns securely generated random bytes.
//
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a securely generated random string.
//
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

var serverSecret string

func ServerSecret() string {
	if serverSecret != "" {
		os.Setenv("SECRET", serverSecret)
		return serverSecret
	}

	serverSecret = GetenvOrSetDefaultFn("SECRET", func() string {
		s, e := GenerateRandomString(32)
		if e != nil {
			panic(e)
		}
		return s
	})
	return serverSecret
}
