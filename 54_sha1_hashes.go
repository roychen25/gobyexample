package gobyexample

import (
	"crypto/sha1"
	"fmt"
)

// SHA1HashesDemo - demonstrates generating SHA1 hashes
func SHA1HashesDemo() {
	s := "sha1 this string"

	// The pattern for generating a hash is `sha1.New()`,
	// `sha1.Write(bytes)`, then `sha1.Sum([]byte{})`.
	// Here we start with a new hash.
	h := sha1.New()

	// `Write` expects bytes. If you have a string, use
	// `[]byte(s)` to coerce it to bytes.
	h.Write([]byte(s))

	// This gets the finalized hash result as a byte slice.
	// The argument to `Sum` can be used to append to an existing
	// byte slice: it usually isn't needed.
	bs := h.Sum(nil)

	// SHA1 values are often printed in hex, e.g. in git commits.
	// Use the %x format verb to convert a hash result into a hex string.
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
