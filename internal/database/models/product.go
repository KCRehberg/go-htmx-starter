package models

// ensure isadmin cant be set by anyone other than an admin
type Product struct {
	Default
	Name   string `json:"name" gorm:"unique;not null"`
	Price  string `json:"price" gorm:"not null"`
	Weight string `json:"weight" gorm:"not null"`
}
