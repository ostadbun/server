package manipulationRepository

import (
	"ostadbun/entity"
)

func (d DB) AddLessonPending(lesson entity.PendingLesson, userId int) error {

	query := `
        insert into pending_lesson(
                                   name ,
                                   difficulty,
                                   description,
                                   name_english,
                                   description_english,
                                   submitted_by
                                   ) values ($1, $2, $3,$4,$5,$6)
    `

	err := d.conn.Conn().QueryRow(query,
		lesson.Name,
		lesson.Difficulty,
		lesson.Description,
		lesson.NameEnglish,
		lesson.DescriptionEnglish,
		userId,
	).Err()

	if err != nil {
		return err
	}

	return nil
}
