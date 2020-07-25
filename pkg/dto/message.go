package dto

// ReplyMarkup ...
type ReplyMarkup struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
}

// Message ...
type Message struct {
	ChatID      string       `json:"chat_id"`
	Text        string       `json:"text"`
	ReplyMarkup *ReplyMarkup `json:"reply_markup"`
}
