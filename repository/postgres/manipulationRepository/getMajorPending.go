package manipulationRepository

import "ostadbun/entity"

// GetMajorPending returns all majors with 'pending' status
func (d DB) GetMajorPending() ([]entity.PendingMajor, error) {
	query := `
        SELECT 
            id,
            name,
            status,
            name_english,
            submitted_by,
            description,
            submitted_at,
            description_english,
            approved_by,
            approved_at,
            rejection_reason
        FROM pending_major
        WHERE status = 'pending'
    `

	rows, err := d.conn.Conn().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var majors []entity.PendingMajor

	for rows.Next() {
		var major entity.PendingMajor
		err := rows.Scan(
			&major.Id,
			&major.Name,
			&major.Status,
			&major.NameEnglish,
			&major.SubmittedBy,
			&major.Description,
			&major.SubmittedAt,
			&major.DescriptionEnglish,
			&major.ApprovedBy,
			&major.ApprovedAt,
			&major.RejectionReason,
		)
		if err != nil {
			return nil, err
		}
		majors = append(majors, major)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return majors, nil
}
