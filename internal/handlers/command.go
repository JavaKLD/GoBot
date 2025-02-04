package handlers

import (
	"GoBotRepo/internal/texts"
	"GoBotRepo/pkg/database"
	"GoBotRepo/pkg/note_model"
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strings"
)

func Start(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   texts.Welcome,
	})
}

func AddNote(ctx context.Context, b *bot.Bot, update *models.Update) {
	userID := update.Message.Chat.ID

	content := strings.TrimSpace(update.Message.Text)

	if content == "" {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: userID,
			Text: "Для создания заметки, отправьте текс заметки",
		})
		return
	}

	dbModel, err := database.LoadDBModel()
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: userID,
			Text: fmt.Sprintf("Ошибка при подключении к бд %v", err),
		})
		return
	}
	dbConn := dbModel.GetConn()

	err = note_model.AddNote(dbConn, userID, content)
	if err != nil {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: userID,
			Text: fmt.Sprintf("Ошибка при добавлении заметки %v", err),
		})
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: userID,
		Text: "Заметка успешно добавлена",
	})


}

func Help(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   texts.Help,
	})
}
