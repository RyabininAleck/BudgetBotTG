package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func Setup() tgbotapi.InlineKeyboardMarkup {
	depositBtn := tgbotapi.NewInlineKeyboardButtonData("Доходы", "Настройки.Доходы")
	wasteBtn := tgbotapi.NewInlineKeyboardButtonData("Расходы", "Настройки.Расходы")

	statusRow := tgbotapi.NewInlineKeyboardRow(depositBtn, wasteBtn)

	currentyBtn := tgbotapi.NewInlineKeyboardButtonData("Валюта", "Настройки.Валюта")
	ReportBtn := tgbotapi.NewInlineKeyboardButtonData("Отчет", "Настройки.Отчет")

	SetupRow := tgbotapi.NewInlineKeyboardRow(currentyBtn, ReportBtn)

	return tgbotapi.NewInlineKeyboardMarkup(statusRow, SetupRow)
}
