package models

// ensure isadmin cant be set by anyone other than an admin
type Account struct {
	Default
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	IsAdmin  bool   `json:"is_admin" gorm:"default:false"`
}
