//閾値3
//次数2(閾値-1)

//pは素数(n<p)

package ShamirSecret

import (
	"encoding/binary"
	"math"
	"math/rand"
)

func new(n int, m string) {
	//setShare
	share := n
	//setSecret
	secret := m

	generatePolynomial(share, secret)
}

func generatePolynomial(n int, m string) []byte {
	//数字として扱う必要があるので文字列をバイト配列に変換

	constant := []byte(m)

	degree := n
	var polynomial float64
	var x float64

	for i := 0; i < degree; i++ {
		float64Degree := float64(i)
		coefficient := rand.Float64()
		//Powがfloat64のみ対応
		polynomial += coefficient * math.Pow(x, float64Degree)
	}
	var bytePolynomial []byte
	//float64をbyteに変換
	bytePolynomial = float64ToByte(polynomial)
	bytePolynomial = append(bytePolynomial, constant...)
	return bytePolynomial
}

func float64ToByte(f float64) []byte {
	var buf []byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}
