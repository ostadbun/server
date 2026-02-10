package entity

type Lesson struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Difficulty  int    `json:"difficulty"`
	Description string `json:"description"`
}
