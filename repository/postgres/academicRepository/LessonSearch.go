package academicRepository

import (
	"ostadbun/entity"
)

func (d DB) LessonSearch(name string) ([]entity.Lesson, error) {
	var lessons []entity.Lesson

	// Query برای جستجوی درس‌ها
	query := `
        SELECT id, name, difficulty, description 
        FROM lesson 
        WHERE name ILIKE '%' || $1 || '%'
    `

	// اجرای Query و دریافت نتایج
	rows, err := d.conn.Conn().Query(query, name)
	if err != nil {
		return nil, err // در صورت خطا، خطا را بازگردانی کن
	}
	defer rows.Close() // بستن نتایج پس از پایان

	// پر کردن لیست درس‌ها
	for rows.Next() {
		var lesson entity.Lesson
		err := rows.Scan(
			&lesson.Id,
			&lesson.Name,
			&lesson.Difficulty,
			&lesson.Description,
		)
		if err != nil {
			return nil, err // در صورت خطا در Scan، خطا را بازگردانی کن
		}
		lessons = append(lessons, lesson)
	}

	// بررسی خطا در حین پیمایش ردیف‌ها
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lessons, nil
}
