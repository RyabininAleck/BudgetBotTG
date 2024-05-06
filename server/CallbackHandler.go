package server

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"BudgetBotTG/models"
	"BudgetBotTG/server/keyboard"
)

func (s *Server) handleCallback(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {

	//todo сделать нормальную дату на английском, желательно вообще json и передавать больше значений
	switch callback.Data {
	case "Настройки.Доходы":
		s.handleSetupCategories(bot, callback, "deposit")
	case "Настройки.Расходы":
		s.handleSetupCategories(bot, callback, "waste")
	case "Настройки.Валюта":
		s.handleCurrency(bot, callback)
	case "Настройки.Отчет":
		s.handleReport(bot, callback)
	case "ДобавитьКатегорию.waste":
		s.handleAddCategory(bot, callback, "waste")
	case "ДобавитьКатегорию.deposit":
		s.handleAddCategory(bot, callback, "deposit")

	case "ДобавитьТранзакцию.waste":
		s.handleAddCategory(bot, callback, "waste")
	case "ДобавитьТранзакцию.deposit":
		s.handleAddCategory(bot, callback, "deposit")

	}

	callbackCfg := tgbotapi.NewCallback(callback.ID, callback.Data)
	if _, err := bot.Request(callbackCfg); err != nil {
		panic(err)
	}
}

func (s *Server) handleAddCategory(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery, status string) {
	// получить данные о категории

	Category := models.Category{
		Name:        "категория1",
		PlannedCost: 100500,
		CurrentCost: 0,
		Status:      status,
	}
	s.storage.PostCategory(callback.From.ID, Category)
	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, callback.Data)
	msg.Text = "категория " + Category.Name + " добавлена"
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}

}

func (s *Server) handleSetupCategories(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery, status string) {
	// todo получить список категорий трат
	// сделать из него кнопочки,
	// отправить
	categories, err := s.storage.GetCategories(callback.From.ID, status)
	if err != nil {
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Ошибка: "+err.Error())
		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, callback.Data)
	msg.ReplyMarkup = keyboard.SetupCategories(categories, status, " Настройки")
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func (s *Server) handleCurrency(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	//todo получить валюту
	// получить список валют

}

func (s *Server) handleReport(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {}
