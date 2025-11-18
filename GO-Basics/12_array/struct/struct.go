package main

// // import (
// // 	"fmt"
// // )

// type Address struct {
// 	City string
// 	Zip  int
// }

// type User struct {
// 	Name string
// 	Addr Address
// }

// type User struct {
// 	Name    string
// 	Age     int
// 	Premium boo
// }
// type Kisi struct {
// 	Isim string
// 	Yas  int
// }

// // func Yenikisi(isim string, yas int) Kisi {
// // 	return Kisi{Isim: isim, Yas: yas}
// // }

// // var u = User{
// // 	Name:    "Nail",
// // 	Age:     25,
// // 	Premium: true,
// // }

// func (u User) Info() {
// 	fmt.Println("User:", u.Name, "Age:", u.Age, "Premium", u.Premium)
// }

// pointer receiver
// struct icinde degisiklik yapmak istiyorsak pointer receiver kullanmaliyiz
// func (u *User) Upgrade() {
// 	u.Premium = false
// }

func main() {
	//basit struct tanimi

	//fmt.Println(User)
	//duz kullanim
	// u := User{"Nail", 25, true}
	// fmt.Println(u)

	//field ismiyle olusturma
	// // u := User{
	// // 	Name:    "Nail",
	// // 	Age:     25,
	// // 	Premium: true,
	// // }
	// // fmt.Println(u)

	//bos struct
	// var u  User
	//---------------------------------------

	//Struct alanlarina erisim
	// u := User{
	// 	Name:    "Nail",
	// 	Age:     25,
	// 	Premium: true,
	// }
	// fmt.Println(u.Name)
	// fmt.Println(u.Age)
	// fmt.Println(u.Premium)
	//---------------------------------

	//STRUCT ILE FONKSIYON

	//-------------------------------------------

	// u.Info()
	// u.Upgrade()
	// -----------------------------------
	// Constructor (yapıcı metod), bir nesne oluşturulduğunda otomatik olarak çalışan özel bir fonksiyondur.
	//
	//	Nesnenin başlangıç durumunu ayarlamak için kullanılır.

	// kisi := Yenikisi("Nail", 30)
	// fmt.Println(kisi.Isim)
	// fmt.Println(kisi.Yas)
	//---------------------------------------

	// u := User{
	// 	Name: "Nail",
	// 	Addr: Address{City: "Siirt", Zip: 56000},
	// }
	// fmt.Println(u)

	//----------------------------------------------------------
	//Slice icinde Struct
	// users := []User{
	// 	{Name: "Ali", Age: 20},
	// 	{Name: "Ayşe", Age: 30},
	// }

	// //map icinde kullnimi
	// userMap := map[string]User{
	// 	"nail": {Name: "Nail", Age: 25},
	// }

	// // Embedding (Go’nun kalıtım karşılığı)
	// type Person struct {
	// 	Name string
	// }

	// type Employee struct {
	// 	Person
	// 	Salary int
	// }

	// e := Employee{
	// 	Person: Person{Name: "Nail"},
	// 	Salary: 5000,
	// }

	// fmt.Println(e.Name) // Person.Name

}
