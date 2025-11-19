package main

import (
	"fmt"
)

// Named return values: Dönüş değerlerine isim veriyorum
// Burada "bolum" ve "kalan" return değişkenleri olarak tanımlanıyor
// Bu yaklaşımın avantajı: Documentation görevi görür, kod kendini açıklar
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
