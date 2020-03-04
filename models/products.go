package models



type Product struct {
	GoodId      int    `json:"goodid"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	PromoMsg    string `json:"promomsg"`
	Image       string `json:"image"`
	Url         string `json:"url"`
	Description string `json:"description"`
	CateId      string `json:"cateid"`
}
