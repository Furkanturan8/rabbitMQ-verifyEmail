package main

import (
	"fmt"
	"log"
	"net/http"
)

// Kullanıcı, onay linkine tıkladığında (örneğin, http://localhost:8080/verify?email=kullanici@example.com),
// bu istek işlenir ve kullanıcıya "E-posta başarıyla doğrulandı" mesajı döner.

func main() {
	// Kullanıcı onayı
	http.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		email := r.URL.Query().Get("email")
		if email == "" {
			http.Error(w, "Geçersiz istek", http.StatusBadRequest)
			return
		}

		// Kullanıcı kaydının onaylandığı simüle ediliyor
		fmt.Fprintf(w, "E-posta başarıyla doğrulandı: %s", email)
		log.Printf("E-posta doğrulandı: %s", email)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
