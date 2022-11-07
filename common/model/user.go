package model

type User struct {
	Username  string `json:"username,omitempty"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      string `json:"role,omitempty"`
}
