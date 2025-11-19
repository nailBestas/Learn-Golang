package main

import (
	"fmt"
)

// kendi kendini cagiran fonksiyonlar
func fibonacci(n int) int {
	//durduran kosul
	//yapmassak sonsuz dongu olusur
	if n <= 1 {
		return n
	}
	//recursive case : kendi kendini cagirma
	//fibanacci tanimi: fib(n) = fib(n-1) + fib(n-2)
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(fibonacci(i), " ") // aralara boşluk ekledik
	}
	fmt.Println() // sonrasında satır sonu ekleyelim
}
