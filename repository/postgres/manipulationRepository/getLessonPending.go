package manipulationRepository

import (
	"ostadbun/entity"
)

// GetLessonPending returns all lessons with 'pending' status
func (d DB) GetLessonPending() ([]entity.PendingLesson, error) {
	query := `
        SELECT 
            id,
            name,
            name_english,
            description_english,
            difficulty,
            description,
            status,
            submitted_by,
            submitted_at,
            approved_by,
            approved_at,
            rejection_reason
        FROM pending_lesson
        WHERE status = 'pending'
    `

	rows, err := d.conn.Conn().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []entity.PendingLesson

	for rows.Next() {
		var lesson entity.PendingLesson
		err := rows.Scan(
			&lesson.Id,
			&lesson.Name,
			&lesson.NameEnglish,
			&lesson.DescriptionEnglish,
			&lesson.Difficulty,
			&lesson.Description,
			&lesson.Status,
			&lesson.SubmittedBy,
			&lesson.SubmittedAt,
			&lesson.ApprovedBy,
			&lesson.ApprovedAt,
			&lesson.RejectionReason,
		)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lessons, nil
}
