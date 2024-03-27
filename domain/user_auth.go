package domain

type UserAuthorizeRequest struct {
	Login    string `json:"login" validate:"required,min=4,max=50"`
	Password string `json:"password" validate:"required"`
}

type UserRegisterRequest struct {
	Login    string `json:"login" validate:"required,min=4,max=50"`
	Password string `json:"password" validate:"required,isDifficultPassword"`
}

type UserRegisterResponse struct {
	ID       uint64 `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type User struct {
	ID       uint64
	Login    string
	Password string
}

type UserDB struct {
	ID       uint64 `db:"id"`
	Login    string `db:"login"`
	Password string `db:"password"`
}
