package academicRepository

import (
	"ostadbun/entity"
)

func (d DB) ProfessorSearch(name string) ([]entity.Professor, error) {
	var professors []entity.Professor

	// Query برای جستجوی اساتید
	query := `
        SELECT id, name, education_history, image_url, description 
        FROM professor 
        WHERE name ILIKE '%' || $1 || '%'
    `

	// اجرای Query و دریافت نتایج
	rows, err := d.conn.Conn().Query(query, name)
	if err != nil {
		return nil, err // در صورت خطا، خطا را بازگردانی کن
	}
	defer rows.Close() // بستن نتایج پس از پایان

	// پر کردن لیست اساتید
	for rows.Next() {
		var professor entity.Professor
		err := rows.Scan(
			&professor.Id,
			&professor.Name,
			&professor.EducationHistory, // فیلد JSONB
			&professor.ImageUrl,
			&professor.Description,
		)
		if err != nil {
			return nil, err // در صورت خطا در Scan، خطا را بازگردانی کن
		}
		professors = append(professors, professor)
	}

	// بررسی خطا در حین پیمایش ردیف‌ها
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return professors, nil
}
