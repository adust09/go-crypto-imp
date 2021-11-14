package main

import (
	"fmt"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
	"crypto/ecdsa"
)

func main() {
	//set a Amount
	sentAmount  := big.NewInt(5)
	receivedAmount  := big.NewInt(5)
	fmt.Println("=================================")
	fmt.Println("Sent Amount: ",sentAmount)
	fmt.Println("Received Amount: ",receivedAmount)
	fmt.Println("=================================")

	//generate a BigNumber
	num1 := GenarateBigNumber()
	num2 := GenarateBigNumber()

	//generate a Keys
	inputKey := hideKey(sentAmount,num1)
	outputKey := hideKey(receivedAmount,num2)

	//calculate the merged key
	addKey := inputKey.Add(inputKey,outputKey)
	mergedKey := num2.Mul(num2,addKey)

	fmt.Println("Merged Key: ",mergedKey)
	fmt.Println("This Key pair is hided")
	fmt.Println("=================================")

	//check input ?= output
	if  receivedAmount.Cmp(sentAmount) == 0 {
		sub := num1.Mul(num1,sentAmount.Sub(sentAmount,receivedAmount))
		mergedAmount := sub.Mul(num1,sub)
		fmt.Println("Merged Amount: ",mergedAmount)
		fmt.Println("This Amount pair is hided")
		fmt.Println("=================================")
	}else {
		fmt.Println("Merged Amount and Merged Key are not equal")
	}
}

func GenarateBigNumber() (*big.Int) {
	//素数生成は時間かかるので注意
	p,e := rand.Prime(rand.Reader, 10000)
	if e != nil {
		fmt.Println("Error: ",e)
	}
	return p
}

func GenarateKey() (*ecdsa.PrivateKey) {
	privateKey,e := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if e != nil {
		fmt.Println("Error: ",e)
	}
	return privateKey
}

func hideAmount(amount,bigNum1 *big.Int) (*big.Int) {
	Amount := bigNum1.Mul(bigNum1,amount)
	return Amount
}

func hideKey(key,bigNum2 *big.Int) (*big.Int) {
	Key := bigNum2.Mul(bigNum2,key)
	return Key
}