package api

import (
	"TGbroadcastservice/internal/telegram"
	"encoding/json"
	"log"
	"net/http"
)

type SendMessageRequest struct {
	Message string `json:"message"`
	UserID  int    `json:"user_id"`
}

func SendMessageHandler(w http.ResponseWriter, r *http.Request, token string) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var req SendMessageRequest

	log.Println("Получен запрос на отправку сообщения")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Ошибка при разборе JSON: %v", err)
		http.Error(w, "Неверный формат запроса: "+err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Запрос: %v", req)

	if err := telegram.SendMessage(token, req.UserID, req.Message); err != nil {
		log.Printf("Ошибка при отправке сообщения: %v", err)
		http.Error(w, "Не удалось отправить сообщение", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Сообщение принято к отправке"))
}
