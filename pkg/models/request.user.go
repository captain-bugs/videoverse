package models

type ReqSaveUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
