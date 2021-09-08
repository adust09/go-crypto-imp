//課題1-1
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	hashFunction := sha256.New()
	hashFunction.Write([]byte("12345"))
	fmt.Printf("%x \n", hashFunction.Sum(nil))

	hashFunction.Write([]byte("12346"))
	fmt.Printf("%x", hashFunction.Sum(nil))
}
