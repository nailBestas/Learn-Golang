package main

import "fmt"

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

// Method (value receiver)
func (u User) Info() {
	fmt.Printf("%s (%d) | Premium: %v | City: %s\n",
		u.Name, u.Age, u.Premium, u.Address.City)
}

// Method (pointer receiver)
func (u *User) Upgrade() {
	u.Premium = true
}

// Constructor
func NewUser(name string, age int, city string) User {
	return User{
		Name:    name,
		Age:     age,
		Premium: false,
		Address: Address{
			City: city,
			Zip:  34000,
		},
	}
}

func main() {

	// Constructor kullanarak user oluşturma
	u1 := NewUser("Nail", 25, "İstanbul")
	u2 := NewUser("Ayşe", 30, "Ankara")

	// Slice içinde struct
	users := []User{u1, u2}

	// Map içinde struct
	mp := map[string]User{
		"nail": u1,
		"ayse": u2,
	}

	// Kullanıcı bilgilerini yazdırma
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
