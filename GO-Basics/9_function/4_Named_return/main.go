package main

import (
	"fmt"
)

func bolmeIslemi(bolunen, bolen int) (bolum int, kalan int) {
	bolum = bolunen / bolen
	kalan = bolunen % bolum

	return
}

func main() {
	fmt.Println("Hello, World!")

	b, k := bolmeIslemi(10, 3)
	fmt.Printf("Bölüm: %d, Kalan: %d\n", b, k)
	fmt.Printf("Bolum: %d, Kalan: %d\n", b, k)
}
