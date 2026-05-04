package users_postgres_repository

type UserModel struct {
	ID           int
	Version      int
	FullName     string
	PasswordHash string
	Email        string
	PhoneNumber  *string
	Telegram     *string
}
