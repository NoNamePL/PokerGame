package models

// Users date

type User struct {
	ID       uint   `json:"id" gorm:"primary key"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type CreateUser struct {
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type UpdateUser struct {
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
