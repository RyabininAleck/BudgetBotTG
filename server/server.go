package server

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"BudgetBotTG/config"
	"BudgetBotTG/storage"
)

type Server struct {
	config  config.ServerConfig
	storage storage.Storage
}

func NewServer(cfg config.ServerConfig, myStorage storage.Storage) *Server {
	err := myStorage.Ping()
	if err != nil {
		return nil
	}
	return &Server{config: cfg, storage: myStorage}
}

func (s *Server) Run() {

	bot, err := tgbotapi.NewBotAPI("6478243484:AAEPfG_z4ASAxZHs2kgtthGeTgxURB0fJ70")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = s.config.Bot.Debug

	u := tgbotapi.NewUpdate(0)
	u.Timeout = s.config.Bot.Timeout
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				s.handleCommand(bot, update.Message)
			}
			if !update.Message.IsCommand() {
				s.handleNotCommand(bot, update.Message)
			}

		} else if update.CallbackQuery != nil {
			s.handleCallback(bot, update.CallbackQuery)
		}
	}

}
