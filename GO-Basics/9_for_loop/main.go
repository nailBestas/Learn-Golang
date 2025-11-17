package main

import (
	"fmt"
)

//c benzeri kullanimi
//	func main() {
//		for i := 0; i < 5; i++ {
//			fmt.Println(i)
//		}
//	}

// 1-while benzeri
// func main() {
// 	i := 0
// 	for i < 5 {
// 		fmt.Println(i)
// 		i++
// 	}
// }

// 2-sonsuz dongu
// Genellikle menü sistemleri veya server loop’larında kullanılır
// func main() {
// 	for {
// 		fmt.Println("Sonsuz dongu")
// 		break
// 	}
// }

// 3-diziler,slice ,map,string
// Iterable Veri Yapısı Nedir?

// Iterate edilebilen (dolaşılabilen) veri yapıları demek.

// Go’da başlıca iterable yapılar:

// Slice → Dinamik diziler ([]int{...})

// Array → Sabit uzunluklu diziler ([5]int{...})

// Map → Key-value veri yapısı (map[string]int{...})

// String → Karakter dizileri (UTF-8)

// range bu yapıların üzerinde index/value veya key/value döndürmek için kullanılır.

// 3.1 slice ornegi
// func main() {
// 	numbers := []int{10, 20, 30}
// 	for i, v := range numbers {
// 		fmt.Println("index :", i, "value:", v)
// 	}

// }

// 3.2 Array ornek --sabit boyutlu
// func main() {
// 	arr := [4]string{"a", "b", "c", "d"}
// 	for i, val := range arr {
// 		fmt.Println(i, val)
// 	}

// }

// 3.3 Map
// func main() {
// 	age := map[string]int{"Nail": 40, "Yusuf": 30}
// 	for name, age := range age {
// 		fmt.Println(name, age)
// 	}
// }

//3.4 String ile

// func main() {
// 	str := "Yusuf"
// 	for i, ch := range str {
// 		fmt.Println(i, string(ch))
// 	}
// }

// Toplu ornek
func main() {
	fmt.Println("Basit for dongusu c tarzinde")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("while benzeri kullanim")
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	fmt.Println("Sonsuz dongu ve break kullanimi")
	break1 := 0
	for {
		if break1 >= 3 {
			break
		}
		fmt.Println(break1)
		break1++
	}

	fmt.Println("Slice uzerinde for + range kullanimi")
	numbers := []int{10, 20, 30}
	for index, value := range numbers {
		fmt.Println(index, value)
	}

	fmt.Println("map uzerinde for kullanimi")
	ages := map[string]int{"Yusuf": 30, "nail": 40}
	for key, val := range ages {
		fmt.Println(key, "yasi", val)
	}

	fmt.Println("contine ornegi")
	for c := 0; c < 5; c++ {
		if c%2 == 0 {
			continue // cift sayilari atla
		}
		fmt.Println(c)
	}
}
