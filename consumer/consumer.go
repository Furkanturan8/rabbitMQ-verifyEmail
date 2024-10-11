package main

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

// Bu aşamada, RabbitMQ kuyruğuna düşen mesajlar (e-posta onay linkleri) alınır ve bir e-posta simülasyonu yapılır.
func main() {
	// RabbitMQ'ya bağlan
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Kanal oluştur
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Kuyruğu al
	q, err := ch.QueueDeclare(
		"email_verification_queue", // kuyruk adı
		true,                       // durable
		false,                      // delete when unused
		false,                      // exclusive
		false,                      // no-wait
		nil,                        // args
	)
	failOnError(err, "Failed to declare a queue")

	// Mesajları al
	msgs, err := ch.Consume(
		q.Name, // kuyruk
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	// Mesajları dinle
	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			log.Printf("Gönderilen Onay Linki: %s", d.Body)
			// Burada e-posta gönderme işlemi yapılabilir.
			// Bir SMTP servisine bağlanarak gerçek e-posta gönderilebilir.
		}
	}()

	log.Printf(" [*] Onay linkleri gönderiliyor. Çıkmak için CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
