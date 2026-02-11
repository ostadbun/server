package manipulationRepository

import (
	"ostadbun/entity"
)

func (d DB) AddMajorPending(major entity.PendingMajor, userId int) error {
	query := `
        INSERT INTO pending_major (
            name, 
            description,
            name_english,
            description_english,
            submitted_by
        ) 
        VALUES ($1, $2, $3, $4, $5)
    `

	err := d.conn.Conn().QueryRow(
		query,
		major.Name,
		major.Description,
		major.NameEnglish,
		major.DescriptionEnglish,
		userId,
	).Err()

	if err != nil {
		return err
	}

	return nil
}
