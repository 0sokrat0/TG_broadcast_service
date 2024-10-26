package main

import (
	"TGbroadcastservice/internal/messaging"
	"TGbroadcastservice/internal/telegram"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {
	// Подключение к NATS с использованием IP контейнера
	natsClient, err := messaging.NewNATSClient("nats://127.0.0.1:4222")
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS: %v", err)
	}
	defer natsClient.Close()

	limiter := time.Tick(33 * time.Millisecond) // Ограничение на 30 сообщений в секунду

	_, err = natsClient.Conn.Subscribe("telegram.send", func(msg *nats.Msg) {
		go func(data []byte) {
			var req telegram.SendMessagePayload
			if err := json.Unmarshal(data, &req); err != nil {
				log.Printf("Ошибка разбора сообщения: %v", err)
				return
			}

			<-limiter

			if err := telegram.SendMessage("YOUR_TELEGRAM_BOT_TOKEN", req.ChatID, req.Text); err != nil {
				log.Printf("Ошибка отправки в Telegram: %v", err)
			} else {
				log.Println("Сообщение успешно отправлено")
			}
		}(msg.Data)
	})

	if err != nil {
		log.Fatalf("Ошибка подписки на тему: %v", err)
	}

	log.Println("Воркер запущен и ждёт сообщений...")
	select {}
}
