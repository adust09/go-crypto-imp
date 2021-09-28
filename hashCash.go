package main

import (
	"crypto/sha256"
	"fmt"
	"bytes"
)

func main(){
	//init difficulty
	h := sha256.New()
	h.Write([]byte("100000"))
	difficulty := h.Sum(nil)
	fmt.Printf("difficulty = %x \n", difficulty)

	//init input
	input := []byte("0")

	for i := 0; i < 100;  {
		h := sha256.New()
		fmt.Printf("index = %x \n", byte(i))
		//create nonce
		h.Write(input)
		result := h.Sum(nil)

		//Compare result with difficulty
		if bytes.Compare(result, difficulty) <= 0 {
			fmt.Printf("founded result is %x \n", result)
			break
		}else{
			i++
			input = append(input, byte(i))
		}
	}
}