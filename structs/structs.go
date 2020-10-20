package structs

import "time"

type ShortURL struct {
	ID          string    `json:"id"`
	OriginalURL string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
	Creator     int       `json:"creator"`
	Anonymous   bool      `json:"anonymous"`
}

type User struct {
	ID               int       `json:"id"`
	Username         string    `json:"username"`
	PasswordHash     string    `json:"password_hash"`
	RegistrationDate time.Time `json:"registration_date"`
	Permission       int       `json:"permission"`
	Email            string    `json:"email"`
}

type APIResponse struct {
	Content interface{} `json:"content"`
	Errors  []string    `json:"errors"`
}
