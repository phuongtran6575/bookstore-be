package models

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
