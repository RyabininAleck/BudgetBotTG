package keyboard

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"BudgetBotTG/models"
)

func SetupCategories(Categories []models.Category, status string, action string) tgbotapi.InlineKeyboardMarkup {

	setupCategoryRow := [][]tgbotapi.InlineKeyboardButton{}

	for _, category := range Categories {
		categoryBtn := tgbotapi.NewInlineKeyboardButtonData(category.Name, action+"."+category.Name)
		categoryRow := tgbotapi.NewInlineKeyboardRow(categoryBtn)
		setupCategoryRow = append(setupCategoryRow, categoryRow)
	}

	addCategoryBtn := tgbotapi.NewInlineKeyboardButtonData("Добавить категорию", "ДобавитьКатегорию."+status)
	addCategoryRow := tgbotapi.NewInlineKeyboardRow(addCategoryBtn)
	setupCategoryRow = append(setupCategoryRow, addCategoryRow)

	return tgbotapi.NewInlineKeyboardMarkup(setupCategoryRow...)
}

func Categories(Categories []models.Category) tgbotapi.InlineKeyboardMarkup {
	CategoryRows := [][]tgbotapi.InlineKeyboardButton{}

	for _, category := range Categories {
		categoryBtn := tgbotapi.NewInlineKeyboardButtonData(category.Name, "ДобавитьТранзакцию."+category.Name)
		categoryRow := tgbotapi.NewInlineKeyboardRow(categoryBtn)
		CategoryRows = append(CategoryRows, categoryRow)
	}

	return tgbotapi.NewInlineKeyboardMarkup(CategoryRows...)
}
