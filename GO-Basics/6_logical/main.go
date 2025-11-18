package main

import (
	"fmt"
)

func main() {
	// a := false
	// b := false
	// fmt.Println(a && b)
	// fmt.Println((a || b))
	// fmt.Println(a != b)

	//KARMA ORNEK OYUNCU ANALIZI EDELIM

	//slice : oyuncularin yaslari
	ages := []int{15, 18, 22, 30, 12, 40}

	//map : kullanici online mi
	online := map[string]bool{
		"nail":  true,
		"yusuf": true,
		"arya":  false,
		"ceylo": false,
		"nihat": true,
	}

	fmt.Println("yas analizi")
	for _, age := range ages {
		if age >= 18 && age <= 30 {
			fmt.Println(age, "genc yetiskin")
		} else if age < 18 {
			fmt.Println(age, "cocuk,genc")

		} else {
			fmt.Println(age, "olgun yetiskin")
		}
	}

	fmt.Println("Online kullanici kontrolu")
	for user, isOnline := range online {
		//mantiksal operatorlerle kullanici online ve ismi 4 harften buyuk olanlari
		if isOnline && len(user) > 4 {
			fmt.Println(user, "online ve uzun isimli kullnicilar")
		}
		//kulllanicilar offline veya ismi 5 harften kucukse
		if !isOnline || len(user) < 5 {
			fmt.Println(user, "offline ve 5 ten kucuk kisler")

		}
	}
	fmt.Println("karisik ornek")
	for _, age := range ages {
		// 18 yas ustu ve (22 yas degil veya 40)
		if age > 18 && (age != 22 || age == 40) {
			fmt.Println(age)
		}

	}

}
