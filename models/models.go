package models

// User представляет модель пользователя.
type User struct {
	UserID      int    `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
	TelegramID  int64  `json:"telegram_id"`
	State       string `json:"state"`
	Language    string `json:"language"`
}

// Category представляет модель категории.
type Category struct {
	CategoryID      int    `json:"category_id"`
	Name            string `json:"name"`
	OwnerTelegramID int64  `json:"owner_telegram_id"`
	PlannedCost     int    `json:"planned_cost"`
	CurrentCost     int    `json:"current_cost"`
	Status          string `json:"category_type"`
}

// Transaction представляет модель транзакции.
type Transaction struct {
	TransactionID   int     `json:"transaction_id"`
	OwnerTelegramID int64   `json:"owner_telegram_idd"`
	Amount          float64 `json:"amount"`
	CategoryID      int     `json:"category_id"`
	Date            string  `json:"date"`
	Comment         string  `json:"comment"`
}
