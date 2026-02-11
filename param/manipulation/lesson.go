package manipulationParam

type PendingLesson struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	Difficulty         int    `json:"difficulty"`
	DescriptionEnglish string `json:"description_english"`
	NameEnglish        string `json:"name_english"`
}
