package types

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"isAdmin"`
}

type UserSlice struct {
	Users []User `json:"users"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
	IsAdmin         bool   `json:"isAdmin"`
}

type Notification struct {
	ProjectID int    `json:"id"`
	Message   string `json:"message"`
}
