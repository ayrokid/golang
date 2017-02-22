package main

import "fmt"

func main() {
	var a, b int

	a = 10
	b = 3

	//jumlah
	c := a + b
	fmt.Printf("%d + %d = %d \n", a, b, c)

	//kurang
	d := a - b
	fmt.Printf("%d - %d = %d \n", a, b, d)

	//bagi
	e := float32(a) / float32(b)
	fmt.Printf("%d / %d = %.2f \n", a, b, e)
}
