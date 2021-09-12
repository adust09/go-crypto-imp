//課題1-2
package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
)

func main() {
	createHashChain()
}

func createHashChain() {
	Array := [...]string{"0001", "0002", "0003", "0004"}
	h := sha256.New()

	//構造体を定義
	type Block struct {
		Index int
		Data  string
		Hash  []byte
	}

	//構造体配列を定義
	type BlockArray []*Block
	var hashChain BlockArray

	for i := 0; i < len(Array); i++ {
		if i == 0 {
			h.Write([]byte(Array[i]))
			hashChain = append(hashChain, &Block{i, Array[i], h.Sum(nil)})
			fmt.Printf("Array %x  = %x \n", i, []byte(Array[i]))
			fmt.Printf("Block %x Hash = %x \n", i, hashChain[i].Hash)
			fmt.Printf("------------------------------------------------------------------------------------- \n")

		} else {

			message := append([]byte(Array[i]), hashChain[i-1].Hash...)
			fmt.Printf("Array %x  = %x \n", i, []byte(Array[i]))
			fmt.Printf("Block %x Hash = %x \n", i-1, hashChain[i-1].Hash)

			h.Write([]byte(message))
			hashChain = append(hashChain, &Block{i, Array[i], h.Sum(nil)})
			fmt.Printf("Block %x Hash = %x \n", i, hashChain[i].Hash)
			fmt.Printf("-------------------------------------------------------------------------------- \n")

		}
	}
	fmt.Printf("================================================================================\n")

	h = sha256.New()

	for i := 0; i < len(Array); i++ {
		fmt.Printf("Array %x  = %x \n", i, []byte(Array[i]))

		if i == 0 {
			h.Write([]byte(Array[i]))
			fmt.Printf("Block %x Hash = %x \n", i, hashChain[i].Hash)
		} else {
			fmt.Printf("Block %x Hash = %x \n", i-1, hashChain[i-1].Hash)
			message := append([]byte(Array[i]), hashChain[i-1].Hash...)
			h.Write([]byte(message))
			fmt.Printf("Block %x Hash = %x \n", i, h.Sum(nil))
		}
		res := h.Sum(nil)

		if reflect.DeepEqual(hashChain[i].Hash, res) {
			fmt.Printf("Block %x is valid \n", i)
		} else {
			fmt.Printf("Block %x is invalid \n", i)
		}
		fmt.Printf("-------------------------------------------------------------------------------- \n")
	}

}
