# GoBot

GoBot - это Telegram-бот, написанный на Go. Он позволяет пользователям хранить, просматривать и удалять заметки.

## Возможности

* Сохранение заметок

* Просмотр сохраненных заметок

* Удаление заметок

## Установка и запуск

### 1. Клонирование репозитория

```bash
git clone https://github.com/JavaKLD/GoBot.git
cd GoBot
```

### 2. Установка зависимостей

```go
go mod tidy
```
### 3. Настройка окружения

Создайте файл .env и укажите в нем:

```
TELEGRAM_BOT_TOKEN=your_token_here
DB_USER=root
DB_PASSWORD=your_password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=gobot
```
### 4. Настройка базы данных MySQL

1. Установите MySQL, если он еще не установлен.

3. Создайте базу данных:
```mysql
CREATE DATABASE gobot;
```
3. Примените миграции (если есть миграционный скрипт).

### 5. Запуск бота

`go run main.go`

## Структура проекта
```
GoBot/
├── cmd/        # Точка входа в приложение
├── internal/   # Основная логика приложения
├── pkg/        # Вспомогательные модули
└── config/     # Конфигурация
```
