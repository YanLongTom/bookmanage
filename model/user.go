package model

type User struct {
	Uid      string `json:"uid"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
