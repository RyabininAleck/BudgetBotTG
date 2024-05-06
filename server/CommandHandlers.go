package server

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"BudgetBotTG/models"
	"BudgetBotTG/server/keyboard"
)

func (s *Server) handleCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	switch message.Command() {

	case "setup":
		s.handleSetup(bot, message)
	case "start":
		s.handleStart(bot, message)

	//todo
	case "report":
		s.handleStart(bot, message)
	case "account":
		s.handleStart(bot, message)
	case "help":
		s.handleStart(bot, message)
	case "export":
		s.handleStart(bot, message)
	case "referal":
		s.handleStart(bot, message)
	}
}

func (s *Server) handleStart(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	lang := message.From.LanguageCode

	msg := tgbotapi.NewMessage(message.Chat.ID, s.config.Chat[lang].HelloMsg)

	u := models.User{
		PhoneNumber: "",
		Name:        message.Chat.UserName,
		TelegramID:  message.From.ID,
		State:       "active",
		Language:    lang,
	}

	s.storage.PostUser(u)

	// todo настройка языка и валюты
	// todo настройка категорий
	// todo перекидывать на настройку категорий
	// msg.ReplyMarkup = setupKeyboard()
	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (s *Server) handleSetup(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "⚙️ Выберите раздел для настройки:")
	msg.ReplyMarkup = keyboard.Setup()
	_, err := bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
