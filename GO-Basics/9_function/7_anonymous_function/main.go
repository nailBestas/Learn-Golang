package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("merhaba ben isimsiz fonksiyon")
	}()

	//degiskene atama
	kareAl := func(x int) int {
		return x * x

	}
	fmt.Println(kareAl(5))

	//closure: Dis scopede ki degiskenelre erisme
	sayac := 0
	arttir := func() int {
		sayac++
		return sayac
	}
	fmt.Println(arttir())
	fmt.Println(arttir())
	fmt.Println(arttir())

}
