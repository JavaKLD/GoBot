package main

import (
	"GoBotRepo/internal/filters"
	"GoBotRepo/internal/handlers"
	"GoBotRepo/internal/texts"
	"GoBotRepo/pkg/systems"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"os"
	"os/signal"
)

func main() {
	token := systems.BotToken()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		panic(err)
	}

	commands := []models.BotCommand{
		{Command: "start", Description: "Начать работу с ботом"},
		{Command: "help", Description: "Получить справочную информацию"},
	}

	comm, _ := b.SetMyCommands(ctx, &bot.SetMyCommandsParams{		//меню помощника по командам
		Commands: commands,
	})
	if !comm {
		panic("Ошибка")
	}

	b.RegisterHandlerMatchFunc(filters.IsStart, handlers.Start)
	b.RegisterHandlerMatchFunc(filters.IsHelp, handlers.Help)

	b.RegisterHandlerMatchFunc(filters.IsPhoto, handlers.Photo)
	b.RegisterHandlerMatchFunc(filters.IsVideo, handlers.Video)
	b.Start(ctx)

}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:  update.Message.Chat.ID,
		Text: texts.Opts,
	})
}