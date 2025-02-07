package handlers

import (
	"GoBotRepo/internal/texts"
	"GoBotRepo/pkg/note_model"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"strings"
)

 var userStates = make(map[int64]string)

const AwaitingNoteCont = "awaiting_note_content"

func Start(ctx context.Context, b *bot.Bot, update *models.Update, db *sql.DB) {
	userId := update.Message.From.ID
	username := update.Message.From.Username

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", userId).Scan(&exists)
	if err != nil {
		log.Println("Ошибка проверки пользователя", err)
		return
	}

	if !exists {
		_, err := db.Exec("INSERT INTO users(id, name) VALUES (?, ?)", userId, username)
		if err != nil {
			log.Println("Ошибка добавления пользователя", err)
			return
		}
		log.Printf("Добавлен пользователь: ID=%d, UserName=%s", userId, username)
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   texts.Welcome,
	})
}

func AddNote(ctx context.Context, b *bot.Bot, update *models.Update, db *sql.DB) {
	userID := update.Message.Chat.ID
	content := strings.TrimSpace(update.Message.Text)

	if content == "/addnote" {
		userStates[userID] = AwaitingNoteCont
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: userID,
			Text: "Отправьте текст",
		})
		return
	}

	if state, exists := userStates[userID]; exists && state == AwaitingNoteCont {
		if content == "" {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: userID,
				Text: "Заметка не может быть пустой",
			})
			return
		}
		err := note_model.AddNote(db, userID, content)
		if err != nil {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: userID,
				Text: fmt.Sprintf("Ошибка при добавлении заметки: %v", err),
			})
			return
		}

		delete(userStates, userID)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: userID,
			Text: "Заметка добавлена",
		})
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
