package entity

type Professor struct {
	Id               string            `json:"id"`
	Name             string            `json:"name"`
	EducationHistory map[string]string `json:"education_history"`
	ImageUrl         string            `json:"image_url"`
	Description      string            `json:"description"`
}
