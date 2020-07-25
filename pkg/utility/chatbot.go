package utility

import (
	"bytes"
	"comet/pkg/dto"
	"encoding/json"
	"net/http"
	"os"
)

// SendBotNotification ...
func SendBotNotification(telegramID string, message string) error {
	telegramMessage := &dto.Message{
		ChatID: telegramID,
		Text:   message,
		ReplyMarkup: &dto.ReplyMarkup{
			RemoveKeyboard: true,
		},
	}

	requestBody, err := json.Marshal(telegramMessage)
	if err != nil {
		return err
	}

	_, err = http.Post("https://api.telegram.org/bot"+os.Getenv("BOT_TOKEN")+"/sendMessage",
		"application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	return nil
}
