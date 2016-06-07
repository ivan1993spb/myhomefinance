package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

// newGUID generates a random GUID
func newGUID() (string, error) {
	guid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, guid)
	if n != len(guid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	guid[8] = guid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	guid[6] = guid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", guid[0:4], guid[4:6], guid[6:8], guid[8:10], guid[10:]), nil
}
