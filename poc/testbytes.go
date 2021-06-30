package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	fmt.Println(math.MaxInt64)

	target := big.NewInt(1)
	fmt.Println(target)
	fmt.Println(uint(256 - 10))
	target.Lsh(target, uint(256-10))

	fmt.Println(target)

	/*b1 := []byte("Ola pequeno mundo bytes")
	b2 := []byte("Ola pequeno mundo 2")

	f := bytes.Join([][]byte{b1, b2}, []byte{})

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(f[:])
	*/
	//fmt.Println(reflect.TypeOf([]byte(teste)))
	//info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
}
