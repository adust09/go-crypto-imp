//課題1-2
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
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
			fmt.Println("index = ", i)

			h.Write([]byte(Array[i]))
			// hashChain = []*Block{
			// 	&Block{i, Array[i], h.Sum(nil)},
			// }
			hashChain = append(hashChain, &Block{i, Array[i], h.Sum(nil)})

			fmt.Printf("Block %x Hash = %x \n", i, hashChain[i].Hash)
		} else {
			fmt.Println("index = ", i)

			message := append([]byte(Array[i]), hashChain[i-1].Hash...)
			h.Write([]byte(message))
			fmt.Printf("Block %x Hash = %x \n", i, h.Sum(nil))

			// hashChain = []*Block{
			// 	&Block{i, Array[i], h.Sum(nil)},
			// }
			hashChain = append(hashChain, &Block{i, Array[i], h.Sum(nil)})

			// fmt.Printf("Block %x Hash = %x \n", i, hashChain[i].Hash)
		}
		fmt.Printf("LatestBlock_Hash= %x \n", hashChain[i].Hash)
	}
}
