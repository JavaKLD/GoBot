package handlers

import (
	"GoBotRepo/internal/texts"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func Photo(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:  update.Message.Chat.ID,
		Text : texts.Photo,
	})
}

func Video(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:  update.Message.Chat.ID,
		Text : texts.Video,
	})
}