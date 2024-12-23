package models

type Product struct {
	ID    int    `gomr:"primaryKey"`
	Name  string `json:"name"`
	Price string `json:"price"`
	Image string `json:"image"`
}
