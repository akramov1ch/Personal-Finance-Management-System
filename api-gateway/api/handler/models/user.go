package models

type RegisterUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type RegisterUserRes struct {
	Message string `json:"message"`
}

type VerifyUserReq struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type VerifyUserRes struct {
	UserId    int    `json:"user_id"`
	Token     string `json:"token"`
	Message   string `json:"message"`
	ExpiresIn int64  `json:"expiresIn"`
}

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRes struct {
	Userid   int    `json:"user_id"`
	Token     string `json:"token"`
	Message   string `json:"message"`
	ExpiresIn int64  `json:"expiresIn"`
}

type ProfileUserRes struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
}
