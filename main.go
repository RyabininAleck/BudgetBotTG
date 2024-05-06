//package main
//
//import (
//	"log"
//
//	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
//)
//
//var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
//	tgbotapi.NewInlineKeyboardRow(
//		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
//		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
//		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
//	),
//	tgbotapi.NewInlineKeyboardRow(
//		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
//		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
//		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
//	),
//)
//
//func main() {
//	bot, err := tgbotapi.NewBotAPI("6478243484:AAEPfG_z4ASAxZHs2kgtthGeTgxURB0fJ70")
//	if err != nil {
//		log.Panic(err)
//	}
//
//	bot.Debug = true
//
//	log.Printf("Authorized on account %s", bot.Self.UserName)
//
//	u := tgbotapi.NewUpdate(0)
//	u.Timeout = 60
//
//	updates := bot.GetUpdatesChan(u)
//
//	// Loop through each update.
//	for update := range updates {
//		// Check if we've gotten a message update.
//		if update.Message != nil {
//			// Construct a new message from the given chat ID and containing
//			// the text that we received.
//			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
//
//			// If the message was open, add a copy of our numeric keyboard.
//			switch update.Message.Text {
//			case "open":
//				msg.ReplyMarkup = numericKeyboard
//
//			}
//
//			// Send the message.
//			if _, err = bot.Send(msg); err != nil {
//				panic(err)
//			}
//		} else if update.CallbackQuery != nil {
//			// Respond to the callback query, telling Telegram to show the user
//			// a message with the data received.
//			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
//			if _, err := bot.Request(callback); err != nil {
//				panic(err)
//			}
//
//			// And finally, send a message containing the data received.
//			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
//			if _, err := bot.Send(msg); err != nil {
//				panic(err)
//			}
//		}
//	}
//}

package main

import (
	"BudgetBotTG/config"
	"BudgetBotTG/middleware"
	"BudgetBotTG/server"
	"BudgetBotTG/storage/sqliteDB"
)

func main() {
	pathToConfig := "" //todo
	serverConfig := config.GetServerConfig(pathToConfig)
	middlewares := middleware.GetMiddleware()
	myStorage := sqliteDB.GetStorage(serverConfig, middlewares)
	myStorage.MakeMigrations()

	srv := server.NewServer(serverConfig, myStorage)

	srv.Run()
	//todo добавить таблицу настроек пользователей.
	// валюта, язык, параметры отчета, категория последнего добавления

	// Добавить таблицу с последним добавлением

}
