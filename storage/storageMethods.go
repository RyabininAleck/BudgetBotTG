package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"BudgetBotTG/middleware"
	"BudgetBotTG/models"
)

// User

func (s *Storage) GetUser(telegramID string) (*models.User, error) {
	middlewares := s.GetMiddlewares()
	middlewares.Start()
	middlewares.Write(middleware.TimeChecker, "GetUser")
	defer middlewares.Stop()

	row := s.conn.QueryRow(selectUser, telegramID)
	var user models.User
	err := row.Scan(&user.UserID, &user.PhoneNumber, &user.Name, &user.TelegramID, &user.State)

	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("Пользователь с указанным ID не найден")
	}

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Storage) PostUser(user models.User) error {
	middlewares := s.GetMiddlewares()
	middlewares.Start()
	middlewares.Write(middleware.TimeChecker, "PostUser")
	defer middlewares.Stop()

	//todo если есть, то не добавлять
	_, err := s.conn.Exec(insertUser, user.PhoneNumber, user.Name, user.TelegramID, user.State) //можно получить result {LastInsertId} id юзера
	if err != nil {
		return err
	}

	return nil

}

// Category

func (s *Storage) GetCategory(telegramID int64, name string) (*models.Category, error) {
	middlewares := s.GetMiddlewares()
	middlewares.Start()
	middlewares.Write(middleware.TimeChecker, "GetCategory")
	defer middlewares.Stop()

	row := s.conn.QueryRow(selectCategory, telegramID, name)
	category := models.Category{Name: name, OwnerTelegramID: telegramID}
	err := row.Scan(&category.CategoryID, &category.Status, &category.PlannedCost, &category.CurrentCost)

	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("Категория с указанным telegramID и name не найдена")
	}

	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (s *Storage) GetCategories(telegramID int64, status string) ([]models.Category, error) {
	middlewares := s.GetMiddlewares()
	middlewares.Start()
	middlewares.Write(middleware.TimeChecker, "GetCategories")
	defer middlewares.Stop()

	rows, err := s.conn.Query(selectCategories, telegramID, status)

	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("Категории с указанным telegramID и status не найдены")
	}
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := []models.Category{}
	for rows.Next() {
		category := models.Category{OwnerTelegramID: telegramID, Status: status}
		err = rows.Scan(&category.CategoryID, &category.Name, &category.PlannedCost, &category.CurrentCost)

		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (s *Storage) UpdateCategory(telegramID int64, oldName string, category models.Category) error {
	middlewares := s.GetMiddlewares()
	middlewares.Start()
	middlewares.Write(middleware.TimeChecker, "UpdateCategory")
	defer middlewares.Stop()

	_, err := s.conn.Exec(updateCategory, category.Status, category.PlannedCost, category.CurrentCost, category.Name, telegramID, oldName)

	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteCategory(telegramID int64, name string) error {
	middlewares := s.GetMiddlewares()
	middlewares.Start()
	middlewares.Write(middleware.TimeChecker, "DeleteCategory")
	defer middlewares.Stop()

	_, err := s.conn.Exec(deleteCategory, telegramID, name)

	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("Категория с указанным telegramID и name не найдена")
	}

	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) PostCategory(telegramID int64, category models.Category) error {
	middlewares := s.GetMiddlewares()
	middlewares.Start()
	middlewares.Write(middleware.TimeChecker, "PostCategory")
	defer middlewares.Stop()

	_, err := s.conn.Exec(insertCategory, telegramID, category.Name, category.PlannedCost, category.CurrentCost, category.Status) //можно получить result {LastInsertId} id юзера
	if err != nil {
		return err
	}

	return nil
}

// Transactions
// надо затестить
func (s *Storage) GetTransactions(telegramID int64, categoryId int, count int) ([]models.Transaction, error) {
	middlewares := s.GetMiddlewares()
	middlewares.Start()
	middlewares.Write(middleware.TimeChecker, "GetCategories")
	defer middlewares.Stop()

	rows, err := s.conn.Query(selectTopTransaction, telegramID, categoryId, count)

	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("Транзакций с указанным telegramID и categoryID не найдены")
	}
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	transactions := []models.Transaction{}
	for rows.Next() {
		transaction := models.Transaction{CategoryID: categoryId}
		err = rows.Scan(&transaction.TransactionID, &transaction.OwnerTelegramID, &transaction.Amount, &transaction.Date, &transaction.Comment) //todo

		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (s *Storage) PostTransactions(transaction models.Transaction) (transactionsID int64, err error) {
	middlewares := s.GetMiddlewares()
	middlewares.Start()
	middlewares.Write(middleware.TimeChecker, "PostTransactions")
	defer middlewares.Stop()

	_, err = s.conn.Exec(insertTransaction, transaction.OwnerTelegramID, transaction.Amount, transaction.CategoryID, transaction.Date, transaction.Comment)
	if err != nil {
		return 0, err
	}

	_, err = s.conn.Exec(updateCategoryCurrentCost, transaction.Amount, transaction.OwnerTelegramID, transaction.CategoryID)
	if err != nil {
		return 0, err
	}

	return transactionsID, nil
}

func (s *Storage) UpdateTransactionsComment(transactionID int64, comment string) error {
	middlewares := s.GetMiddlewares()
	middlewares.Start()
	middlewares.Write(middleware.TimeChecker, "UpdateTransactionsComment")
	defer middlewares.Stop()

	_, err := s.conn.Exec(updateTransactionComment, comment, transactionID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) UpdateTransactionsData(transactionID int64, date time.Time) error {
	middlewares := s.GetMiddlewares()
	middlewares.Start()
	middlewares.Write(middleware.TimeChecker, "UpdateTransactionsData")
	defer middlewares.Stop()

	_, err := s.conn.Exec(updateTransactionData, date, transactionID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) DeleteTransactions(transactionID int64) error {
	middlewares := s.GetMiddlewares()
	middlewares.Start()
	middlewares.Write(middleware.TimeChecker, "DeleteTransactions")
	defer middlewares.Stop()

	_, err := s.conn.Exec(deleteTransaction, transactionID)
	if err != nil {
		return err
	}
	return nil
}
