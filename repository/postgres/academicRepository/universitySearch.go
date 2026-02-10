package academicRepository

import (
	"ostadbun/entity"
)

func (d DB) UniversitySearch(name string) ([]entity.University, error) {
	var universities []entity.University

	// Query برای جستجوی دانشگاه‌ها
	query := `
        SELECT id, name, city, category, image_url, description 
        FROM university 
        WHERE name ILIKE '%' || $1 || '%'
    `

	// اجرای Query و دریافت نتایج
	rows, err := d.conn.Conn().Query(query, name)
	if err != nil {
		return nil, err // در صورت خطا، خطا را بازگردانی کن
	}
	defer rows.Close() // بستن نتایج پس از پایان

	// پر کردن لیست دانشگاه‌ها
	for rows.Next() {
		var university entity.University
		err := rows.Scan(
			&university.Id,
			&university.Name,
			&university.City,
			&university.Category,
			&university.ImageUrl,
			&university.Description,
		)
		if err != nil {
			return nil, err // در صورت خطا در Scan، خطا را بازگردانی کن
		}
		universities = append(universities, university)
	}

	// بررسی خطا در حین پیمایش ردیف‌ها
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return universities, nil
}
