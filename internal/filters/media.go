package filters

import "github.com/go-telegram/bot/models"

func IsPhoto(update *models.Update) bool {
	if update.Message == nil {
		return false
	}

	return len(update.Message.Photo) != 0
}

func IsVideo(update *models.Update) bool {
	if update.Message == nil {
		return false
	}

	return update.Message.Video != nil
 }
