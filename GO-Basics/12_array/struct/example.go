package main

import (
	"fmt"
)

type Address struct {
	City string
	Zip  int
}
type User struct {
	Name    string
	Age     int
	Premium bool
	Address Address
}

// method (value receiver) ir struca bagli fonksiyonlardir
func (u User) Info() {
	fmt.Println("%s (%d) | Premium: %v | City: %s\n", u.Name, u.Age, u.Premium, u.Address.City)
}

// pointer receiver
func (u *User) Upgrade() {
	u.Premium = true
}

// constructor
func NewUser(name string, age int, city string) User {
	return User{
		Name:    name,
		Age:     age,
		Premium: false,
		Address: Address{
			City: city,
			Zip:  56000,
		},
	}
}
func main() {

	//constructor kullanarak user olusturma
	u1 := NewUser("Nail", 40, "siirt")
	u2 := NewUser("Ayse", 38, "Sori")

	//slice icinde struct
	users := []User{u1, u2}

	//Map icinde struct
	mp := map[string]User{
		"nail": u1,
		"ayse": u2,
	}

	//Kullanici bilgilerini yazdirma
	for _, user := range users {
		user.Info()
	}
	fmt.Println("\nNail Premium yapılıyor...")
	u1.Upgrade()
	u1.Info()

	fmt.Println("\nMap'ten kullanıcı okutma:")
	if user, ok := mp["nail"]; ok {
		user.Info()
	}

}
