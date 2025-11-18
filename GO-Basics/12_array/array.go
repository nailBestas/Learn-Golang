package main

import (
	"fmt"
)

//	func change(arr *[3]int) {
//		arr[0] = 999
//	}
func main() {
	// 	//array
	// 	var numbers [3]int
	// 	fmt.Println(numbers)

	// 	//deger atama
	// 	numbers[0] = 10
	// 	numbers[1] = 11
	// 	numbers[2] = 12
	// 	fmt.Println(numbers)

	// 	//arrayi direkt tanimlama-sabit boyutlu bir dizi
	// 	numbers2 := [5]int{13, 14, 15, 16, 17}
	// 	fmt.Println(numbers2)

	// 	//otomatik boyut ayarlama-dinamik boyutlu ve eleman erisme
	// 	numbers3 := []int{1, 2, 3, 4, 5, 6}
	// 	fmt.Println(numbers3)
	// 	fmt.Println(numbers3[3])

	// 	//array uzerinde dongu
	// 	for i := 0; i < len(numbers3); i++ {
	// 		fmt.Println(numbers3[i])
	// 		//fmt.Println(numbers3[i])
	// 	}

	// 	//range ile
	// 	for index, val := range numbers3 {
	// 		fmt.Println(index, val)
	// 	}

	// 	//string olarak array
	// 	items := []string{"a", "b", "c", "d"}
	// 	fmt.Println(items)

	// 	//Array hafiza da calisma mantigi

	// 	//Array kopyalanir-referans verilmez-bir veri yapisidir
	// 	a := [3]int{1, 2, 3}
	// 	b := a
	// 	b[0] = 10
	// 	fmt.Println(a)
	// 	fmt.Println(b)

	// 	//Array fonksiyona gonderildiginde de kopyalanir
	// 	x := [3]int{6, 7, 8}
	// 	change(&x)
	// 	fmt.Println(x)

	// 	//Array boyutu tipe  dahil degildir
	//uzunluklari veya boyutlri farkli oldugu icin farkli tur aliyoruz
	// var a [5]int
	// var b [4]int
	// fmt.Println(a = b)

	//arrayda hafiza tek bloktur go da cok hizli sesbi budur
	// // arr := [3]int{1,2,3}

	//range ile kopyalama durumu
	// burada value degeri yani v degeri array eleamaninin kopyasidir
	// arr := [3]int{1, 2, 3}
	// for i, v := range arr {
	// 	fmt.Println(i, v)
	// }

	//ornek
	// x := [3]int{1, 2, 3}
	// y := x
	// y[1] = 6
	// fmt.Println("x degeri", x, "y degeri :", y)

	//COK BOYUTLU DIZILER 2D VE 3D
	// var matrix [2][3]int{

	// }
	// 2D ARRAY
	// numbers := [2][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// }
	// fmt.Println(numbers)
	// //dis for dongusu satirlar uzerinde
	// for rowIndex, row := range numbers {
	// 	//ic for sutunlar uzerinde yapilir
	// 	for colIndex, val := range row {
	// 		fmt.Println(rowIndex, colIndex, val)
	// 	}
	// }

	// 3D ARRAY(YUKSELIK X SATIRX SUTUN)
	// var cube [2][2][2]int

	// 3D array ornek
	cube := [2][2][2]int{
		{
			{1, 2},
			{3, 4},
		},
		{
			{5, 6},
			{7, 8},
		},
	}
	fmt.Println(cube)
	for i, layer := range cube {
		for j, row := range layer {
			for k, value := range row {
				fmt.Println(i, j, k, value)
			}
		}
	}
}
