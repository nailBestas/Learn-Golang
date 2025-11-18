package main

import (
	"fmt"
)

func main() {
	// arr := [5]int{1, 2, 3, 4, 5}
	// s := arr[1:5]
	// fmt.Println(s)
	//

	// 3 yontemle slice tanimlariz
	// s := []int{1, 2, 3}
	// s := make([]int, 3)
	// s ;= make([]int,3 ,10) //len 3 cap 10

	//append ile ekleme
	// arr := []int{1, 2, 3, 4, 5} //Eğer capacity doluysa → Go yeni bir array oluşturur Eski verileri taşıyıp slice pointer’ını yeni array’e bağlar
	// arr = append(arr, 6)
	// fmt.Println(arr)

	//copy
	// dst := make([]int, 3)
	// src := []int{10, 20, 30, 40, 50}
	// n := copy(dst, src)
	// fmt.Println(n)

	//slice Append davranis
	// arr := [5]int{1, 2, 3, 4, 5}
	// // 0 index dahil ama end index dahil degil
	// s := arr[1:4]
	// fmt.Println(s)
	// fmt.Println(len(arr))
	// fmt.Println(cap(arr))
	// fmt.Println(len(s))
	// fmt.Println(cap(s))
	// //kapasite yeterli oldu yeni array olusmadi
	// // s hala eski arraya bagli
	// s = append(s, 99)
	// fmt.Println(s)

	// kapasite yetmnezse ne olacak
	// s := []int{1, 2}
	// s2 := s
	// s = append(s, 3, 4, 5, 6) //kapasite yetersiz -yeni array olusacak
	// fmt.Println(s)
	// fmt.Println(s2)
	//birbirinden farkli tamamen ayri 2 farkli array olustu

	// slice trap
	// s := []int{1, 2, 3, 4}
	// s2 := s[:2]
	// s = append(s, 99)
	// fmt.Println(s2)
	// fmt.Println(s)

	//slice icerigini tamamen kopyalamak
	// s := []int{1, 2, 3, 4}
	// s2 := make([]int, len(s))
	// s3 := copy(s, s2)
	// fmt.Println(s3)

	//SLICE PROFESYONEL ORNEK

	//baslangic slice
	users := []string{"nail", "ceylin", "arya"}
	fmt.Println("users", users)

	//capacity control
	fmt.Println("len :", len(users), "cap :", cap(users))

	//kapasiteyi buyutelim
	users = append(users, "ayse", "steve")
	fmt.Println("son eklenenler : ", users)

	//slice kopyalama
	backup := make([]string, len(users))
	copy(backup, users)

	//orjinal degistirelim
	users[0] = "degisti"

	fmt.Println("users :", users)
	fmt.Println("backup: ", backup)

}
