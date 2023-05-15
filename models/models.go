package models

import "time"

type Customer struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName string `gorm:"not null" json:"firstname"`
	LastName  string `gorm:"not null" json:"lastname"`
	Email     string `gorm:"not null;unique" json:"email"`
	Phone     string `json:"phone"`
	Avatar    string `json:"avatar"`
}

type Product struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Supplier struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Category struct {
	ID   uint `json:"id"`
	Name uint `json:"name"`
}

type Order struct {
	ID        uint      `json:"id"`
	OrderDate time.Time `json:"orderdate"`
}

type Cart struct {
	ID         uint `json:"id"`
	TotalItems uint `json:"totalitems"`
}

type Review struct {
	ID      uint   `json:"id"`
	Rating  uint   `json:"rating"`
	Heading string `json:"heading"`
	Text    string `json:"text"`
}
