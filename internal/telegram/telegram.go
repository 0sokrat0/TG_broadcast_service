package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Структура для отправки сообщения
type SendMessagePayload struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}

type SendPhotoPayload struct {
	ChatID  int    `json:"chat_id"`
	Photo   string `json:"photo"`
	Caption string `json:"caption"`
}

type SendDocumentPayload struct {
	ChatID   int    `json:"chat_id"`
	Document string `json:"document"`
	Caption  string `json:"caption"`
}

func SendPhoto(token string, chatID int, photoURL, caption string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", token)

	payload := SendPhotoPayload{
		ChatID:  chatID,
		Photo:   photoURL,
		Caption: caption,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("ошибка маршалинга JSON: %w ", err)
	}
	return sendRequest(url, body)
}

func SendDocument(token string, chatID int, documentURL, caption string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendDocument", token)
	payload := SendDocumentPayload{
		ChatID:   chatID,
		Document: documentURL,
		Caption:  caption,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("ошибка маршалинга JSON: %w ", err)
	}
	return sendRequest(url, body)
}

func SendMessage(token string, chatID int, message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	payload := map[string]interface{}{
		"chat_id": chatID,
		"text":    message,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("ошибка маршалинга JSON: %w", err)
	}

	return sendRequest(url, body)
}

func sendRequest(url string, body []byte) error {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("ошибка создания запроса: %w", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("ошибка от Telegram API: статус %d", resp.StatusCode)
	}

	return nil
}
