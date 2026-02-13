package manipulationRepository

import (
	"ostadbun/entity"
)

// GetUniversityPending returns all universities with 'pending' status
func (d DB) GetUniversityPending() ([]entity.PendingUniversity, error) {
	query := `
        SELECT 
            id,
            name,
            name_english,
            description_english,
            city,
            category,
            image_url,
            description,
            status,
            submitted_by,
            submitted_at,
            approved_by,
            approved_at,
            rejection_reason
        FROM pending_university
        WHERE status = 'pending'
    `

	rows, err := d.conn.Conn().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var universities []entity.PendingUniversity

	for rows.Next() {
		var university entity.PendingUniversity
		err := rows.Scan(
			&university.Id,
			&university.Name,
			&university.NameEnglish,
			&university.DescriptionEnglish,
			&university.City,
			&university.Category,
			&university.ImageUrl,
			&university.Description,
			&university.Status,
			&university.SubmittedBy,
			&university.SubmittedAt,
			&university.ApprovedBy,
			&university.ApprovedAt,
			&university.RejectionReason,
		)
		if err != nil {
			return nil, err
		}
		universities = append(universities, university)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return universities, nil
}
