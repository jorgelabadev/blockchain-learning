package main

import (
	"bytes"
	"fmt"
)

func main() {

	b1 := []byte("Ola pequeno mundo bytes")
	b2 := []byte("Ola pequeno mundo 2")

	f := bytes.Join([][]byte{b1, b2}, []byte{})

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(f[:])

	//fmt.Println(reflect.TypeOf([]byte(teste)))
	//info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
}
