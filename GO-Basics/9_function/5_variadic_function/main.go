package main

import (
	"fmt"
)

func sayilariTopla(sayilar ...int) int {
	fmt.Printf("sayilar tipi: %T\n", sayilar)
	toplam := 0

	for _, sayi := range sayilar {
		toplam += sayi //toplam = toplam + sayi
	}

	return toplam
}

func main() {

	fmt.Println(sayilariTopla(1, 2))
	fmt.Println(sayilariTopla(1, 2, 3))
	fmt.Println(sayilariTopla(1, 2, 3, 4))

	sayilar := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sayilariTopla(sayilar...))

}
