package storage

// User
var (
	// insertUser добавляет пользователя phone_number, name, telegram_id, state
	insertUser = "INSERT INTO users (phone_number, name, telegram_id, state) VALUES (?, ?, ?, ?);"

	// selectUser получает одного пользователя по telegram id
	selectUser = `
		SELECT user_id, phone_number, name, telegram_id, state
		FROM users
		WHERE telegram_id = ?
	`
)

// Category
var (
	// insertCategory добавляет категорию для пользователя: telegram_id, name, planned_cost,current_cost, status
	insertCategory = `INSERT INTO categories (owner_telegram_id, name, planned_cost, current_cost, status)
						VALUES (?, ?, ?, ?, ?)`

	// deleteCategory удаляет категорию по telegram_id, name
	deleteCategory = `
		DELETE FROM categories
		WHERE
			owner_telegram_id = ? AND name = ?
	`

	// selectCategory получить категорию пользователя по telegram id и name
	selectCategory = `
		SELECT category_id, status, planned_cost , current_cost
		FROM categories 
		WHERE owner_telegram_id = ? AND name = ?
	`

	// updateCategory обновленине категории по telegram id и name
	updateCategory = `	UPDATE categories
	SET status = ?, planned_cost = ?, current_cost = ?, name = ?
	WHERE owner_telegram_id = ? AND name = ?`

	// selectCategory получить категории пользователя по telegram id и status
	selectCategories = `
		SELECT category_id, name , planned_cost , current_cost
		FROM categories 
		WHERE owner_telegram_id = ? AND status = ?`
)

// Transaction
var (
	// insertTransaction
	insertTransaction = `
	INSERT INTO transactions (owner_telegram_id, amount, category_id, date, comment)
	VALUES (?, ?, ?, ?, ?)`

	// updateCategoryCurrentCost
	updateCategoryCurrentCost = `
	UPDATE categories
	SET current_cost = current_cost + ?
	WHERE category_id = ?`

	updateTransactionComment = `
        UPDATE transactions
        SET comment = ?
        WHERE transaction_id = ?`
	updateTransactionData = `
        UPDATE transactions
        SET date = ?
        WHERE transaction_id = ? `

	deleteTransaction = `
        DELETE FROM transactions
        WHERE transaction_id = ?`

	selectTopTransaction = `
		SELECT transaction_id, owner_telegram_id, amount, date, comment
		FROM transactions
		WHERE owner_telegram_id = ?
		  AND category_id = ?
		ORDER BY date DESC
		LIMIT ?;
		`
)
