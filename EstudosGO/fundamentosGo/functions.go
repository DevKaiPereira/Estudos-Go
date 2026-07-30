package main

func somar(a, b int) int {
	return a + b
}

func higerOrder(c int) func(int) int {
	return func(d int) int {
		return c + d
	}
}

func swap(a, b int) (int, int) {
	return b, a
}

func dividir(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return res, rem
}

func SomarTres(nums ...int) int {
	var out int
	for _, n := range nums {
		out += n
	}
	return out
}
