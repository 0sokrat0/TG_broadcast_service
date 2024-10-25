package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SendMessagePayload struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}

func SendMessage(token string, chatID int, message string) error {
	// Формируем URL для отправки сообщения
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	// Подготавливаем JSON-пейлоад
	payload := SendMessagePayload{
		ChatID: chatID,
		Text:   message,
	}

	// Преобразуем структуру в JSON
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("ошибка преобразования в JSON: %w", err)
	}

	// Выполняем POST-запрос к Telegram API
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("ошибка при отправке запроса: %w", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		// Читаем тело ответа для отладки
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("ошибка от Telegram API: статус %d, ответ: %s", resp.StatusCode, string(respBody))
	}

	return nil // Сообщение успешно отправлено
}
