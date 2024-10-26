package api

import (
	"TGbroadcastservice/internal/telegram"
	"encoding/json"
	"log"
	"net/http"
)

type SendMessageRequest struct {
	Message   string `json:"message"`
	UserID    int    `json:"user_id"`
	MediaType string `json:"media_type"`
	MediaURL  string `json:"media_url"`
}

func SendMessageHandler(token string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}

		var req SendMessageRequest

		// Читаем и парсим JSON после декодирования
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Неверный формат запроса: "+err.Error(), http.StatusBadRequest)
			return
		}

		// Логируем сообщение для отладки
		log.Printf("Отправляется сообщение: %s", req.Message)

		var err error
		if req.MediaType == "photo" {
			err = telegram.SendPhoto(token, req.UserID, req.MediaURL, req.Message)
		} else if req.MediaType == "document" {
			err = telegram.SendDocument(token, req.UserID, req.MediaURL, req.Message)
		} else {
			err = telegram.SendMessage(token, req.UserID, req.Message)
		}

		if err != nil {
			log.Printf("Ошибка при отправке: %v", err)
			http.Error(w, "Не удалось отправить сообщение", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Сообщение принято к отправке"))
	}
}
