package main

import "fmt"

func main() {
	a := 10.0
	b := 3.0

	// toplama
	sum := a + b
	fmt.Println(sum)

	//cikarma
	diff := a - b
	fmt.Println(diff)

	//carpma
	mul := a * b
	fmt.Println(mul)

	//Bolme
	var div float64 = a / b
	fmt.Println(div)

	//Modul
	//mod := a % b
	//fmt.Println(mod)

	//Artirma
	a++
	fmt.Println(a)

	//azalma
	b--
	fmt.Println(b)

	//int8 turu -128 ile 127
	//overflow tasma

	var x int8 = 127
	fmt.Println("baslangic x :", x)

	x = x + 1
	fmt.Println("X in yeni degeri :", x)

	//underflow alt tasma degeri

	var y int8 = -128
	fmt.Println("ynin baslangic degeri : ", y)

	y = y - 1
	fmt.Println("y nin yeni degeri :", y)

}
