package main

import (
	"fmt"
)

type MatematikFonksiyonu func(int, int) int

// normal bir toplama fonksiyomu
func topla(a, b int) int {
	return a + b
}

func carp(a, b int) int {
	return a * b
}
func islemYap(a, b int, f MatematikFonksiyonu) int {
	return f(a, b)
}

func main() {

	var islem MatematikFonksiyonu
	islem = topla
	fmt.Println(islem(5, 3))

	islem = carp
	fmt.Println(islem(5, 3))

	//functionu parametre olarak gonderelim
	sonuc := islemYap(5, 3, topla)
	fmt.Println(sonuc)

	sonuc = islemYap(5, 3, carp)
	fmt.Println(sonuc)

}
