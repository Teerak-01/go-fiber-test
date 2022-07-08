package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

type Employee struct {
	gorm.Model
	EmployeeID int       `json:"employee_id"`
	Name       string    `json:"name"`
	LName      string    `json:"lname"`
	Birthday   time.Time `json:"birthday"`
	Age        int       `json:"age"`
	Email      string    `json:"email"`
	Tel        string    `json:"tel"`
}
