package manipulationRepository

import (
	"ostadbun/entity"
)

func (d DB) AddUniversityPending(university entity.PendingUniversity, userId int) error {
	query := `
        INSERT INTO pending_university (
            name, 
            city, 
            category,
            image_url,
            description,
            name_english,
            description_english,
            submitted_by
        ) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `

	err := d.conn.Conn().QueryRow(
		query,
		university.Name,
		university.City,
		university.Category,
		university.ImageUrl,
		university.Description,
		university.NameEnglish,
		university.DescriptionEnglish,
		userId,
	).Err()

	if err != nil {
		return err
	}

	return nil
}
