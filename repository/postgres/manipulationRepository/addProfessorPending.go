package manipulationRepository

import (
	"ostadbun/entity"
)

func (d DB) AddProfessorPending(professor entity.PendingProfessor, userId int) error {
	query := `
        INSERT INTO pending_professor (
            name, 
            education_history, 
            image_url,
            description,
            name_english,
            description_english,
            submitted_by
        ) 
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `

	err := d.conn.Conn().QueryRow(
		query,
		professor.Name,
		professor.EducationHistory, // باید []byte یا string باشد (json.Marshal شده)
		professor.ImageUrl,
		professor.Description,
		professor.NameEnglish,
		professor.DescriptionEnglish,
		userId,
	).Err()

	if err != nil {
		return err
	}

	return nil
}
