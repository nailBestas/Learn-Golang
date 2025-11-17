package main

import "fmt"

func main() {

	user := "nail"
	password := "1234"
	balance := 450.0
	isBlocked := false
	roles := []string{"user", "premium"}
	settings := map[string]bool{"twoFA": true}

	// --- SHORT STATEMENT + IF ---
	if loginUser := authenticate(user, password); loginUser {

		// --- GUARD PATTERN ---
		if isBlocked {
			fmt.Println("Hesabınız bloke edilmiş!")
			return
		}

		fmt.Println("Giriş başarılı!")

		// --- IC ICE IF (nested if) ---
		if settings["twoFA"] {
			fmt.Println("2FA doğrulandı.")

			if balance > 400 {
				fmt.Println("Hesap bakiyesi:", balance)

				// --- SLICE içinde arama + IF ---
				for _, role := range roles {
					if role == "premium" {
						fmt.Println("Premium özelliklere erişim açıldı.")
					}
				}

			} else {
				fmt.Println("Bakiye yetersiz!")
			}
		}

		// --- ELSE IF ZİNCİRİ ---
		if balance >= 1000 {
			fmt.Println("Altın seviye kullanıcı!")
		} else if balance >= 500 {
			fmt.Println("Gümüş seviye kullanıcı!")
		} else if balance >= 100 {
			fmt.Println("Bronz seviye kullanıcı!")
		} else {
			fmt.Println("Yeni kullanıcı!")
		}

	} else {
		fmt.Println("Hatalı kullanıcı adı veya şifre!")
	}
}

func authenticate(u, p string) bool {
	return u == "nail" && p == "1234"
}
