//Bu satır Go programının bir ana paket (main package) olduğunu gösterir.

// main paketi, çalıştırılabilir bir programın başlangıç noktasıdır.
package main

import (
	"fmt"
)

//Burada yeni bir veri tipi (struct) oluşturuyoruz: Dikdortgen.
// struct birden fazla alanı (property) bir arada tutan bir veri yapısıdır.

// Dikdortgen’in iki alanı var:

// en → dikdörtgenin eni (int tipinde)

// boy → dikdörtgenin boyu (int tipinde)
type Dikdortgen struct {
	en, boy int
}

// method
// Metod, bir struct ile ilişkili fonksiyondur.
// Yani normal fonksiyon gibi çalışır ama o struct’a ait bir özellikmiş gibi davranır.
// RECEIVER'ın detayları:
// (d diktortgen) = receiver
// d = receiver değişkeni (istediğimiz ismi verebiliriz)
// Dikdortgen = receiver'ın tipi
func (x Dikdortgen) Alan() int {
	return x.en * x.boy
}

func main() {
	d := Dikdortgen{en: 3, boy: 4}

	// Method expression
	alanFonk := Dikdortgen.Alan
	fmt.Println(alanFonk(d))
}
