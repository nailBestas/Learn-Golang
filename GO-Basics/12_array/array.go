package main

import (
	"fmt"
)

func main() {
	//array
	var numbers [3]int
	fmt.Println(numbers)

	//deger atama
	numbers[0] = 10
	numbers[1] = 11
	numbers[2] = 12
	fmt.Println(numbers)

	//arrayi direkt tanimlama-sabit boyutlu bir dizi
	numbers2 := [5]int{13, 14, 15, 16, 17}
	fmt.Println(numbers2)

	//otomatik boyut ayarlama-dinamik boyutlu ve eleman erisme
	numbers3 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(numbers3)
	fmt.Println(numbers3[3])

	//array uzerinde dongu
	for i := 0; i < len(numbers3); i++ {
		fmt.Println(numbers3[i])
		//fmt.Println(numbers3[i])
	}

	//range ile
	for index, val := range numbers3 {
		fmt.Println(index, val)
	}

	//string olarak array
	items := []string{"a", "b", "c", "d"}
	fmt.Println(items)

	//
}
