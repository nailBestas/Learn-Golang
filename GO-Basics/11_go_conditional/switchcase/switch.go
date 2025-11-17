package main

import (
	"fmt"
)

func main() {
	role := "admin"      //ornek rol
	balance := 750.0     //hesap bakiyesi
	action := "withdraw" //kullanciinij yapmak istedgi islem

	fmt.Println("kullnici rol kontrolu")

	//1- basic switch (role gore yapalim)
	switch role {
	case "admin":
		fmt.Println("admin paneline hos geldiniz")
	case "premium":
		fmt.Println("premium ozelliler aktif")
	case "user":
		fmt.Println("standar kullnici giris")
	default:
		fmt.Println("bilinmeyen yol")
	}

	fmt.Println("coklu  case ornegi")
	//2-multply case
	switch role {
	case "admin", "moderator", "owner":
		fmt.Println("yonetim yetkileri tanimli")
	default:
		fmt.Println("yonetim yetkisi yok")
	}

	fmt.Println("Sartli switch if-else zinzirine alternatif")
	//3-sart switch
	switch {
	case balance >= 1000:
		fmt.Println("Altin seviyesinde hesap")
	case balance >= 500:
		fmt.Println("gumus seviyesinde hesap")
	case balance >= 100:
		fmt.Println("bronz seviyesi hesap")
	default:
		fmt.Println("yeni hesap")
	}

	fmt.Println("islem tipip")
	//4-Switch ile islem yonlendirme
	switch action {
	case "deposit":
		fmt.Println("para yatirma islemi yapiliyor....")
	case "withdraw":
		if balance > 0 {
			fmt.Println("para cekme islemi yapiliyor...")
		} else {
			fmt.Println("bakiye yok")
		}
	case "transfer":
		fmt.Println("transfer islemi baslatildi...")
	default:
		fmt.Println("gecersiz islem")
	}

	fmt.Println("type swicth ornegi")
	//5-type switch interface ile birlikte calisiyor
	var x interface{} = 25.5
	switch v := x.(type) {
	case int:
		fmt.Println("x bir integer", v)
	case float64:
		fmt.Println("x bir float64", v)
	case string:
		fmt.Println("x bir string", v)
	default:
		fmt.Println("x bilinmeyen tur")
	}
}
