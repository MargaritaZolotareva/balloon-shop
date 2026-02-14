package api

type SimilarProductResponse struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Photo string  `json:"photo_path"`
}
type PhotoResponse struct {
	ID        int    `json:"id"`
	PhotoPath string `json:"photo_path"`
}

type ProductsSectionResponse struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Price         float64 `json:"price"`
	Photo         string  `json:"photo"`
	CategoryTitle string  `json:"category_title"`
}

type GetProductsByCategoryResponse struct {
	CategoryTitle string                    `json:"category_title"`
	Products      []ProductsSectionResponse `json:"products"`
}

type ProductResponse struct {
	ID              int                      `json:"id"`
	Title           string                   `json:"title"`
	Description     string                   `json:"description"`
	Price           float64                  `json:"price"`
	Category        int                      `json:"category"`
	Photos          []PhotoResponse          `json:"photos"`
	SimilarProducts []SimilarProductResponse `json:"similarProducts"`
}

type CategoriesResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Photo string `json:"photo"`
}
