package main

import (
	"fmt"
)

// defer function return olmadan hemen once calistirilacak statementler
// LIFO yapisinda calisir - ilk giren son cikar
func dosyaIslemler() {
	fmt.Println("dosya aciliyor")

	//bu satir function return etmenden once calisacak
	defer fmt.Println("Dosya kapatiliyor....")

	//
	defer fmt.Println("temizlik yapiliyor...")

	fmt.Println("dosya veri yaziliyor.....")
}
func main() {
	dosyaIslemler()
}
