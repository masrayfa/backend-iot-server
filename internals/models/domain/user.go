package domain

type User struct {
	IdUser   int64
	Email    string
	Username string
	Password string
	Status   bool
	Token    string
	IsAdmin  bool
}