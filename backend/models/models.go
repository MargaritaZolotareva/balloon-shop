package models

import "time"

type Product struct {
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Photos       []Photo `gorm:"foreignkey:ProductID"`
	CategoryId   int     `json:"category"`
	ProductOrder int     `json:"order"`
}

type Photo struct {
	ID        int       `json:"id"`
	ProductID int       `json:"product_id"`
	PhotoPath string    `json:"photo_path"`
	CreatedAt time.Time `json:"created_at"`
	PicOrder  int       `json:"pic_order"`
}

type LeadForm struct {
	ID      int    `json:"id"`
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Message string `json:"message"`
}

type Category struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
