package db

import "github.com/3nt3/urlshortener/structs"

func CreateUser(user structs.User) (int, error) {
	const statement = "INSERT INTO users (id, username, email, password_hash, permission, registration_date) VALUES (default, $1, $2, $3, $4, $5) RETURNING id;"
	var id int
	err := database.QueryRow(statement, user.Username, user.Email, user.PasswordHash, user.Permission, user.RegistrationDate).Scan(&id)
	return id, err
}

func GetUserById(id int) (structs.User, error) {
	const query = "SELECT (id, username, email, password_hash, permission, registration_date) FROM users WHERE id = $1;"
	var user structs.User
	err := database.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Permission, &user.RegistrationDate)
	return user, err
}

func GetUserByUsername(username string) (structs.User, error) {
	const query = "SELECT (id, username, email, password_hash, permission, registration_date) FROM users WHERE username = $1;"
	var user structs.User
	err := database.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Permission, &user.RegistrationDate)
	return user, err
}
