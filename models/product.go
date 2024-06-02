package models

import "time"

type Product struct {
	ID        string     `json:"id" gorm:"primary_key;not null;type:varchar(100);index"`
	Name      string     `json:"name" gorm:"not null;type:varchar(250);index"`
	Price     float64    `json:"price" gorm:"index"`
	Stock     float64    `json:"stock" gorm:"index"`
	Status    Status     `json:"status" gorm:"not null;type:varchar(10);index"`
	CreatedAt time.Time  `json:"created_at" gorm:"not null;default:now()"`
	CreatedBy string     `json:"created_by" gorm:"not null;type:varchar(150)"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"default:null"`
	UpdatedBy *string    `json:"updated_by,omitempty" gorm:"type:varchar(150);default:null"`
}

func (Product) TableName() string {
	return "product"
}

type ProductRegister struct {
	Name   string  `json:"name" binding:"required,min=3"`
	Price  float64 `json:"price" binding:"required"`
	Stock  float64 `json:"stock" binding:"required"`
	Status Status  `json:"status"`
}

type ListProduct struct {
	Page      int       `json:"page"`
	Limit     int       `json:"limit"`
	Total     int       `json:"total"`
	TotalPage int       `json:"totalPage"`
	Products  []Product `json:"products"`
}

type ProductUpdate struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Stock     float64 `json:"stock"`
	Status    Status  `json:"status"`
	UpdatedBy string  `json:"updated_by"`
}
