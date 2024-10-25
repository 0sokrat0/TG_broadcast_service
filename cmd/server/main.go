package main

import (
	"TGbroadcastservice/internal/api"
	"TGbroadcastservice/internal/config"
	"fmt"
	"log"
	"net/http"
)

func main() {

	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Ошибка при загрузке конфигурации: %v", err)
	}

	// Регистрируем обработчик и передаем токен с помощью замыкания
	http.HandleFunc("/send-message", func(w http.ResponseWriter, r *http.Request) {
		api.SendMessageHandler(
			w,
			r,
			cfg.Telegram.Token,
		)
	})

	// Запускаем HTTP-сервер на указанном порту
	port := "8080"
	log.Printf("Сервер запущен на http://localhost:%s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
