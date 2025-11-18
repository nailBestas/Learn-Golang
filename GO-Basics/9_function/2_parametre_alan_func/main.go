package main

import (
	"fmt"
)

func kisiSelamla(isim string) string {

	return "Merhaba" + " " + isim + "!"

}

func main() {
	fmt.Println("Hello, World!")

	sonuc := kisiSelamla("Arya")
	fmt.Println(sonuc)
}
