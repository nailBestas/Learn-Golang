package main

import (
	"fmt"
)

func main() {
	//1.map olusturma yontemi
	// ages := map[string]int{
	// 	"nail": 40,
	// 	"ayse": 35,
	// }
	//ages["ceylo"] = 3
	//fmt.Println(ages)
	//fmt.Println(ages["nail"])
	//silme islemi
	//delete(ages, "nail")

	//fmt.Println(len(ages))

	//make ile map olusturma
	//bos map olusturma
	//scores := make(map[string]int)
	//fmt.Println(scores)
	//map anahtar deger girme value atama
	//scores["go"] = 1
	//fmt.Println(scores)

	//nil map
	//nil icine eleman eklenemez
	//var m map[string]int
	// m["x"] = 12
	// fmt.Println(m)
	//---------------------------------------------------
	// map ten veri cekme
	// value int tipi degeri arar
	//exist ise bool turu
	// value, exists := ages["nail"]
	// if exists {
	// 	fmt.Println("buldum", value)
	// } else {
	// 	fmt.Println("boyle bir anahatar yok")
	// }
	//---------------------------------------------------
	//map + range (iteration)
	//anahtar deger
	// for anahtar, deger := range ages {
	// 	fmt.Println(anahtar, deger)
	// }

	// 	//key degerini istemedigmiz zaman
	// for _, age := range ages {
	// 	fmt.Println(age)
	// }

	// //value degeri istemesek
	// for name := range ages {
	// 	fmt.Println(name)
	// }

	//map icindeki elemalar asla ayni sirada donmez
	//------------------------------------------------------
	//map + literal slice icinde
	// students := map[string][]int{
	// 	"Arya":  {100, 90},
	// 	"ceylo": {90, 100},
	// }
	// fmt.Println(students)
	//-------------------------------------------------
	//Map icinde map
	// permission := map[string]map[string]bool{
	// 	"admin": {
	// 		"read":  true,
	// 		"write": true,
	// 	},
	// 	"user": {
	// 		"read":  true,
	// 		"write": false,
	// 	},
	// }
	// fmt.Println(permission["admin"]["read"])
	//-------------------------------------
	// m1 := map[string]int{"x": 1}
	// m2 := m1
	// m2["x"] = 5
	// fmt.Println(m2["x"])
	// fmt.Println(m1["x"])
	//map değişkeni = sadece pointer içerir
	//kopyalanınca aynı tabloyu işaret eder

	//---------------------------------
	//deep copy icin cozum
	// m1 := map[string]int{"a": 10, "b": 20, "c": 30}
	// m2 := make(map[string]int)

	// for k, v := range m1 {
	// 	m2[k] = v
	// }
	// fmt.Println(m2)
	//------------------------------------------

	//profesyonel bir ornek
	users := map[string]map[string]interface{}{
		"nail": {
			"age":     25,
			"premium": true,
			"roles":   []string{"user", "admin"},
		},
		"ayse": {
			"age":     30,
			"premium": false,
			"roles":   []string{"user"},
		},
	}

	// --- Eleman okuma + comma-ok kontrolü ---
	if data, ok := users["nail"]; ok {
		fmt.Println("Nail bulundu:", data)
	}

	// --- Eleman ekleme ---
	users["mehmet"] = map[string]interface{}{
		"age":     40,
		"premium": true,
		"roles":   []string{"user", "editor"},
	}

	// --- Silme ---
	delete(users, "ayse")

	// --- Map içinde döngü ---
	for username, info := range users {
		fmt.Println("Kullanıcı:", username)

		for key, value := range info {
			fmt.Printf("  %s: %v\n", key, value)
		}
	}

	// --- Deep copy örneği ---
	backup := make(map[string]map[string]interface{})
	for k, v := range users {
		backup[k] = v // Bu hala shallow copy — tam deep copy değil!
	}
}
