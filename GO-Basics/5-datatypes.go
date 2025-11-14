package main

import (
	"fmt"
)

// Veri tip olusturmak icin type ile basliyoruz ondan sonraki degisken ismi buyuk harfle basliyor
// Struct tipi
type Kisi struct {
	Ad    string
	Yas   int
	Aktif bool
}

//“Hesaplayici adında yeni bir veri tipi oluşturdum.
// Bu veri tipi, 2 int alan ve int döndüren bir fonksiyondur.”

//Bunu şu şekilde düşün:
//Hesaplayici → fonksiyon tipidir
//Yani Go’da fonksiyon bile bir veri tipi olabilir.

type Hesaplayici func(int, int) int

func main() {

	//Temel veri tipleri

	var yas int = 30
	var isim string = "Nafiz"
	var maas float64 = 67500.75
	var aktif bool = true

	fmt.Println("Yas :", yas)
	fmt.Println("Isim :", isim)
	fmt.Println("Maas:", maas)
	fmt.Println("Aktif :", aktif)

	// Array -sabir liste anlaminda

	sayilar := [3]int{10, 20, 30}
	fmt.Println("Array:", sayilar)

	// Slice -dinamik liste
	diller := []string{"python", "Go", "Rust"}
	fmt.Println("Slice:", diller)

	// Map
	notlar := map[string]int{
		"Nafiz": 92,
		"Nail":  90,
	}
	notlar["Nafiz"] = 100
	fmt.Println("Map : ", notlar)

	// Struct
	kisi1 := Kisi{
		Ad:    "Nafiz",
		Yas:   18,
		Aktif: false,
	}
	fmt.Println("Struct :", kisi1)

	// Pointer
	x := 50
	p := &x
	fmt.Println("Pointer :", p)
	fmt.Println("p nin gosterdigi deger :", *p)

	*p = 100
	fmt.Println("x in yeni degeri :", x)

	// Fonksiyon tipi
	//fonksyionlari degiskenler turmak
	//carpma adında bir değişken oluştur
	// türü → Hesaplayici olsun
	//içeriği → bir fonksiyon olsun
	var carp Hesaplayici = func(a int, b int) int {
		return a * b

	}
	sonuc := carp(10, 20)
	fmt.Println("Fonksiyon tipi sonucu: ", sonuc)

	// fonksioyn ismi icin farkli bir ornek
	//Fonksiyon tipini tanımlayalım:

	type Islem func(int, int) int

	//İki farklı fonksiyon:

	carp1 := func(a, b int) int { return a * b }
	cikar := func(a, b int) int { return a - b }

	//Hepsi Islem tipinde olabilir:

	var i1 Islem = carp1
	var i2 Islem = cikar

	fmt.Println(i1(3, 4))  // 12
	fmt.Println(i2(10, 3)) // 7



	// Normal fonksiyon
func topla(a, b int) int {
    return a + b
}

// Fonksiyon değişkene atanmış
topla := func(a, b int) int {
    return a + b
}

//Fonksiyon tipi ile atanmış
type Hesaplayici func(int, int) int

var topla Hesaplayici = func(a, b int) int {
    return a + b
}


//Üçü de aynı işi yapar.

}
