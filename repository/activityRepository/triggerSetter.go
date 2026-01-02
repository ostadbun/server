package activityRepository

import (
	"fmt"
	"ostadbun/pkg/constants"
)

func (d DB) TriggerSetter(userid int, activity Activityconstants.ActivityTriggersName) error {
	// کوئری را در یک خط ترکیب می‌کنیم:
	// آی‌دی مربوط به آن activity را پیدا کن و در جدول history درج کن
	query := `
        INSERT INTO activity_history (user_id, activity_id)
        SELECT $1, id FROM activity WHERE name = $2
    `

	// چون این یک دستور نوشتن (INSERT) است، از Exec استفاده می‌کنیم
	_, err := d.conn.Conn().Exec(query, userid, activity)

	fmt.Println(activity)
	// حتماً باید ارور را برگردانید تا تابع فراخوان بداند مشکل پیش آمده یا نه
	if err != nil {
		return err
	}

	return nil
}
