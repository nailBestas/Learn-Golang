package main

import (
	"fmt"
)

func hesapMakinesi(a int, b int) (int, int) {

	toplam := a + b
	fark := a - b

	return toplam, fark

}

func main() {
	fmt.Println("Hello, World!")

	t, f := hesapMakinesi(10, 5)
	fmt.Println(t, f)
}
