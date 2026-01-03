package models

import "time"

type Product struct {
	ID           int64   `json:"id"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	Photos       []Photo `gorm:"foreignkey:ProductID"`
	CategoryId   int     `json:"category"`
	ProductOrder int     `json:"order"`
}

type Photo struct {
	ID        int64     `json:"id"`
	ProductID int64     `json:"product_id"`
	PhotoPath string    `json:"photo_path"`
	CreatedAt time.Time `json:"created_at"`
	PicOrder  int64     `json:"pic_order"`
}

type LeadForm struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

type Category struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}
