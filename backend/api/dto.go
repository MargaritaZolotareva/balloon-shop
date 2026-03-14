package api

type SimilarProductResponse struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Slug  string  `json:"slug"`
	Photo string  `json:"photo_path"`
}
type PhotoResponse struct {
	ID        int    `json:"id"`
	PhotoPath string `json:"photo_path"`
}

type HomepageCategory struct {
	Title    string                    `json:"title"`
	Slug     string                    `json:"slug"`
	Products []ProductsSectionResponse `json:"products"`
}

type HomepageResponse struct {
	Categories []HomepageCategory `json:"categories"`
}

type ProductsSectionResponse struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Price         float64 `json:"price"`
	Photo         string  `json:"photo"`
	Slug          string  `json:"slug"`
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
	Slug            string                   `json:"slug"`
	Photos          []PhotoResponse          `json:"photos"`
	SimilarProducts []SimilarProductResponse `json:"similarProducts"`
}

type CategoriesResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Photo string `json:"photo"`
	Slug  string `json:"slug"`
}
