package db

func initURLsTable() error {
	statement := "CREATE TABLE IF NOT EXISTS urls (id VARCHAR(16) PRIMARY KEY, original_url text, created_at timestamp);"
	_, err := database.Exec(statement)
	return err
}

func initUsersTable() error {
	statement := "CREATE TABLE IF NOT EXISTS users (id int PRIMARY KEY, username text, email text, password_hash VARCHAR(128), permission int, registration_date timestamp);"
	_, err := database.Exec(statement)
	return err
}
