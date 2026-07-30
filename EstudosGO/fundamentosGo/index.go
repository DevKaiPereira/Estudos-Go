package main

import  (
	"fmt""fmt"
)

func index()  {
	fmt.Println("Hello, World!")

	fmt.Println(pacotes.Foo)
	pacotes.PrintMinha()

	fmt.Println(somar(1, 2))
	x := higerOrder(2)(1)
	fmt.Println(x)
	fmt.Println(SomarTres(10, 10, 10))

	a, b := swap(10, 20)
	fmt.Println(a, b)
	res, rem := dividir(10, 3)
	fmt.Println(res, rem)
}
