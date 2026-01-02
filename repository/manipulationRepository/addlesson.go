package manipulationRepository

import (
	"fmt"
	manipulationParam "ostadbun/param/manipulation"
)

func (d DB) AddLesson(lesson manipulationParam.Lesson, userId int) error {

	fmt.Println(lesson, userId, "aw3233")

	query := `
        insert into lesson(name , difficulty, description,name_english,description_english,registered_by) values ($1, $2, $3,$4,$5,$6)
    `

	err := d.conn.Conn().QueryRow(query, lesson.Name, lesson.Difficulty, lesson.Description, lesson.NameEnglish, lesson.DescriptionEnglish, userId).Err()

	if err != nil {
		return err
	}

	return nil
}
