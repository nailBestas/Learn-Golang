package main

import (
	"fmt"
	"math"
	"time"
)

//Sabitler blogu

const (
	PI      = 3.14             //tip belirtilmemis
	AppName = "Hesap Makinesi" //string constant
	MaxTry  = 3                // integer constant

	// iota ornegi //enum gibi
	AddOperation = iota //0
	SubOperation        //1
	MulOperation
	DivOperation
)

func main() {
	fmt.Println("Uygulama adi :", AppName)
	fmt.Println("PI:", PI)
	fmt.Println("maksimum deneme:", MaxTry)

	fmt.Println("\nislem kodlari (iota ornegi): ")
	fmt.Println("Toplama: ", AddOperation)
	fmt.Println("cikarma: ", SubOperation)
	fmt.Println("carpma: ", MulOperation)
	fmt.Println("bolme: ", DivOperation)

}
