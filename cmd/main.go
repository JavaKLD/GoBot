package main

import (
	"GoBotRepo/internal/filters"
	"GoBotRepo/internal/handlers"
	"GoBotRepo/internal/texts"
	"GoBotRepo/pkg/database"
	"GoBotRepo/pkg/systems"
	"context"
	"database/sql"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"log"
	"os"
	"os/signal"
)

var DB *sql.DB

func main() {

	dbModel, err := database.LoadDBModel()
	if err != nil {
		log.Fatal("Ошибка загрузки бд %v", err)
	}

	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbModel.User, dbModel.Password, dbModel.Host, dbModel.Port, dbModel.Name)

	DB, err = sql.Open("mysql", dbConn)
	if err != nil {
		log.Fatal("Ошибка при подключении к бд %v", err)
	}
	defer DB.Close()

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

	comm, _ := b.SetMyCommands(ctx, &bot.SetMyCommandsParams{ //меню помощника по командам
		Commands: commands,
	})
	if !comm {
		panic("Ошибка")
	}

	b.RegisterHandlerMatchFunc(filters.IsStart, handlers.Start)
	b.RegisterHandlerMatchFunc(filters.IsHelp, handlers.Help)

	b.Start(ctx)

}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   texts.Opts,
	})
}
