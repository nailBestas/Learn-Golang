package main

import "fmt"

func bankaParaCek(bakiye, cekilecek float64) {
	fmt.Printf("🏦 PARA ÇEKME İŞLEMİ: Bakiye: %.2f TL, Çekilecek: %.2f TL\n", bakiye, cekilecek)

	//  DEFER 1: İşlem günlüğü (NE OLURSA OLSUN kayıt tut)
	defer fmt.Println(" İşlem günlüğüne kaydedildi")

	//  DEFER 2: Panic yakalayıcı (KRİTİK - recover burada)
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf(" ACİL DURUM: %v\n", r)
			fmt.Println(" Sistem güvende, işleme devam ediliyor...")
		}
	}()

	//  DEFER 3: Bakiye kontrolü (işlem sonrası)
	defer func() {
		fmt.Printf(" Son bakiye: %.2f TL\n", bakiye)
	}()

	//  Güvenlik kontrolü - CRITICAL HATA
	if cekilecek > 10000 {
		panic(" Güvenlik ihlali! 10.000 TL üstü çekim yapılamaz!")
	}

	//  Bakiye kontrolü - CRITICAL HATA
	if cekilecek > bakiye {
		panic(" Yetersiz bakiye! Hesapta yeterli para yok!")
	}

	//  Normal işlem
	bakiye -= cekilecek
	fmt.Printf(" Başarılı! %.2f TL çekildi\n", cekilecek)

	//  DEFER 4: Başarı mesajı (en son çalışacak)
	defer fmt.Println("🎊 İşlem tamamlandı!")
}

func main() {
	fmt.Println("========== BANKA SİSTEMİ ==========")

	//  DURUM 1: Normal işlem
	fmt.Println("\n1. NORMAL İŞLEM:")
	bankaParaCek(5000, 1000)

	//  DURUM 2: Yetersiz bakiye - PANIC!
	fmt.Println("\n2. YETERSİZ BAKİYE:")
	bankaParaCek(500, 1000) // PANIC ama YAKALANACAK!

	//  DURUM 3: Güvenlik ihlali - PANIC!
	fmt.Println("\n3. GÜVENLİK İHLALİ:")
	bankaParaCek(50000, 15000) // PANIC ama YAKALANACAK!

	//  DURUM 4: Başka normal işlem
	fmt.Println("\n4. YENİ NORMAL İŞLEM:")
	bankaParaCek(3000, 500)

	fmt.Println("\n========== SİSTEM ÇALIŞMAYA DEVAM EDİYOR ==========")
}
