package filters

import "github.com/go-telegram/bot/models"

const (
	start   = "/start"
	addnote = "/addnote"
	help    = "/help"
)

func IsStart(update *models.Update) bool {
	return update.Message != nil && update.Message.Text == start
}

func IsAdd(update *models.Update) bool {
	return update.Message != nil && update.Message.Text == addnote
}

func IsHelp(update *models.Update) bool {
	return update.Message != nil && update.Message.Text == help
}
