package main

import (
	"TGbroadcastservice/internal/api"
	"TGbroadcastservice/internal/config"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Загружаем конфигурацию
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("Ошибка при загрузке конфигурации: %v", err)
	}

	// Регистрируем обработчик с токеном через замыкание
	http.HandleFunc("/send-message", api.SendMessageHandler(cfg.Telegram.Token))

	// Запускаем HTTP-сервер на указанном порту
	port := "8080"
	log.Printf("Сервер запущен на http://localhost:%s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
	}
}
