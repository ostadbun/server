package manipulationParam

type PendingProfessor struct {
	Name               string            `json:"name"`
	EducationHistory   map[string]string `json:"education_history"`
	ImageUrl           string            `json:"image_url"`
	Description        string            `json:"description"`
	NameEnglish        string            `json:"name_english"`
	DescriptionEnglish string            `json:"description_english"`
}
