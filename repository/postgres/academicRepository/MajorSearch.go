package academicRepository

import (
	"ostadbun/entity"
)

func (d DB) MajorSearch(name string) ([]entity.Major, error) {
	var majors []entity.Major

	// Query برای جستجوی رشته‌ها
	query := `
        SELECT id, name 
        FROM major 
        WHERE name ILIKE '%' || $1 || '%'
    `

	// اجرای Query و دریافت نتایج
	rows, err := d.conn.Conn().Query(query, name)
	if err != nil {
		return nil, err // در صورت خطا، خطا را بازگردانی کن
	}
	defer rows.Close() // بستن نتایج پس از پایان

	// پر کردن لیست رشته‌ها
	for rows.Next() {
		var major entity.Major
		err := rows.Scan(
			&major.Id,
			&major.Name,
		)
		if err != nil {
			return nil, err // در صورت خطا در Scan، خطا را بازگردانی کن
		}
		majors = append(majors, major)
	}

	// بررسی خطا در حین پیمایش ردیف‌ها
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return majors, nil
}
