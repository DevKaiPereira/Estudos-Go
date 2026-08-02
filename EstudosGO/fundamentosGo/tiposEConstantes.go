package main

import "fmt"

//bool

// int int8 int16 int32 int64 (São numero inteiros possitivos e negativos)
// uint uint8 uint16 uint32 uint64(São numeros somente negativos)

//byte

//rune

//float32 float64

//complex64 complex128

//string

func main() {
	var b uint8 = 10
	takeByte(b)
}

func takeByte(b byte) {
	fmt.Println(b)
}

func x() {
	x := 64
	s := string(x)

	fmt.Println(s)
}
