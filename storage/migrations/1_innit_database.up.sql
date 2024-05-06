CREATE TABLE IF NOT EXISTS users (
                                     user_id INTEGER PRIMARY KEY AUTOINCREMENT,
                                     phone_number TEXT,
                                     name TEXT,
                                     telegram_id INTEGER,
                                     state TEXT  CHECK (state IN ('active', 'inactive', 'pending')),
                                     UNIQUE (telegram_id)
);

CREATE TABLE IF NOT EXISTS categories (
                                          category_id INTEGER PRIMARY KEY AUTOINCREMENT,
                                          name TEXT,
                                          owner_telegram_id INTEGER,
                                          planned_cost INTEGER,
                                          current_cost INTEGER,
                                          status TEXT CHECK (status IN ('deposit', 'waste')),
                                          FOREIGN KEY(owner_telegram_id) REFERENCES users(telegram_id),
                                          UNIQUE (name, owner_telegram_id)
);

CREATE TABLE IF NOT EXISTS transactions (
                                            transaction_id INTEGER PRIMARY KEY AUTOINCREMENT,
                                            owner_telegram_id INTEGER,
                                            amount INTEGER,
                                            category_id INTEGER,
                                            date TEXT,
                                            comment TEXT,
                                            FOREIGN KEY(owner_telegram_id) REFERENCES users(telegram_id),
                                            FOREIGN KEY(category_id) REFERENCES categories(category_id)
);

