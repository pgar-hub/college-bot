package photo

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func GetPhoto() *tgbotapi.FileBytes {
	photoBytes, err := os.ReadFile("photo/photo.jpeg")
	if err != nil {
		log.Println("Error reading photo:", err)
		return nil
	}

	photoFile := tgbotapi.FileBytes{
		Name:  "photo.jpg",
		Bytes: photoBytes,
	}
	return &photoFile

}
