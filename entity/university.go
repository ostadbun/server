package entity

type University struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Category    string `json:"category"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
}
