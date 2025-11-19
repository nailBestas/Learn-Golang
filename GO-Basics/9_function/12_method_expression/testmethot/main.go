package main

import "fmt"

type Telefon struct {
	marka string
	model string
}

// SENİN GÖREVİN: Telefon'a method'lar yazmak!
// Aşağıdaki method'ları tamamla:

// Method 1: Telefonu açar
func (t Telefon) ac() {
	fmt.Println(t.marka + " " + t.model + " açılıyor...")
}

// Method 2: Telefonu kapatır
// BURAYI SEN TAMAMLA:
func (t Telefon) kapat() {
	// Buraya kodu yaz
	fmt.Println(t.marka + " " + t.model + "kapaniyor.....")
}

// Method 3: Mesaj gönderir
// BURAYI SEN TAMAMLA:
func (t Telefon) mesajGonder(kisi string) {
	// Buraya kodu yaz - "Ali'ye mesaj gönderiliyor..." gibi
	kisi = "Nail"
	fmt.Println(kisi + "adli kisi " + t.marka + "  " + t.model + " telefonu ile mesaj gonderdii ")
}

// Method 4: Fotoğraf çeker
// BURAYI SEN TAMAMLA:
func (t Telefon) fotografCek() {
	// Buraya kodu yaz
	fmt.Println(" Nail" + t.marka + " ve" + t.model + " telefonu ile" + " " + "Fotograf cekti")
}

func main() {
	// Telefon oluştur
	iphone := Telefon{marka: "iPhone", model: "15 Pro"}

	// Telefon'a işler yaptıralım:
	iphone.ac()
	iphone.mesajGonder("Nail")
	iphone.fotografCek()
	iphone.kapat()
}
