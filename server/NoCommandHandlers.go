package server

import (
	"fmt"

	"github.com/Knetic/govaluate"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"BudgetBotTG/models"
	"BudgetBotTG/server/keyboard"
)

// handleNotCommand обработка сообщения о трате или доходе
func (s *Server) handleNotCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {

	tx, err := parseTransactionMsg(message)
	if err != nil {
		msg := tgbotapi.NewMessage(message.Chat.ID, "Не понимаю. Это не запись")
		if _, err = bot.Send(msg); err != nil {
			panic(err)
		}
		return
	}
	s.handleTransaction(bot, message, tx)

}

func (s *Server) handleTransaction(bot *tgbotapi.BotAPI, message *tgbotapi.Message, tx models.Transaction) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "транзакция: "+fmt.Sprintf("%v", tx)+"\nВыберите категорию:")

	status := "waste"

	categories, err := s.storage.GetCategories(message.From.ID, status)
	if err != nil {
		return
	}
	msg.ReplyMarkup = keyboard.SetupCategories(categories, status, "Добавить")

	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}

}

func parseTransactionMsg(message *tgbotapi.Message) (models.Transaction, error) {

	eval, err := govaluate.NewEvaluableExpression(message.Text)
	if err != nil {
		return models.Transaction{}, err
	}

	result, err := eval.Evaluate(nil)
	if err != nil {
		return models.Transaction{}, err
	}

	return models.Transaction{
		TransactionID:   0,
		OwnerTelegramID: 0,
		Amount:          result.(float64),
		CategoryID:      0,
		Date:            message.Time().String(),
		Comment:         "",
	}, nil
}
