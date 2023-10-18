package util

import "crypto/rand"

func ReadRandom(size int) ([]byte, error) {
	buf := make([]byte, size)
	_, err := rand.Read(buf)

	return buf, err
}
