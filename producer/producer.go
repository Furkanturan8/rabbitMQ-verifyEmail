package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
)

// http://localhost:8081/register?email=kullanici@example.com kullanıcı buraya get isteği gönderdiğinde RabbitMQ kuyruğuna bir onay mesajı eklenir.
// Bu mesaj, kullanıcının onaylaması gereken bir bağlantıdır.

func main() {
	// RabbitMQ'ya bağlan
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Kanal oluştur
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Kuyruk oluştur (onay mesajları için)
	q, err := ch.QueueDeclare(
		"email_verification_queue", // Kuyruk ismi
		true,                       // durable (kalıcı kuyruk)
		false,                      // delete when unused
		false,                      // exclusive
		false,                      // no-wait
		nil,                        // args
	)
	failOnError(err, "Failed to declare a queue")

	// Kayıt fonksiyonu
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		// Kullanıcı bilgilerini al (örnek: email)
		email := r.URL.Query().Get("email")
		if email == "" {
			http.Error(w, "Email is required", http.StatusBadRequest)
			return
		}

		// Onay linki oluştur
		verificationLink := fmt.Sprintf("http://localhost:8080/verify?email=%s", email)

		// Mesajı kuyrukta yayınla
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key (kuyruk adı)
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(verificationLink),
			})
		failOnError(err, "Failed to publish a message")

		log.Printf("Onay linki gönderildi: %s", verificationLink)
		fmt.Fprintf(w, "Kayıt işlemi başarılı, lütfen e-postanızı kontrol edin.")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
