package main

import (
	"errors"
	"fmt"
)

//go da error handling: multiple return values ile deger dondurme
// bu pattern go da standarttit.(result,error ) seklinde

func divide(a, b float64) (float64, error) {
	//hata kontrolu yapalim
	if b == 0 {
		//error new ile yeni bir hata olusturalim
		//iki tarafta da 0 olmali
		return 0.0, errors.New("sifira bolme hatasi")
	}
	return a / b, nil //nil hata yok demek
}

func main() {
	// hata kontrolu yapmadan once
	sonuc, hata := divide(10.0, 2.0)
	if hata != nil {
		fmt.Println(hata)
	} else {
		fmt.Println(sonuc)
	}

	//hata durumunu test edelim
	sonuc, hata = divide(10.0, 0.0)
	if hata != nil {
		fmt.Println(hata)
	} else {
		fmt.Println(sonuc)
	}

}
