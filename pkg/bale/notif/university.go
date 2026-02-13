package notify

import (
	"fmt"
	"ostadbun/entity"
)

// UniversityNotificationData contains university information for notification
type UniversityNotificationData struct {
	Name         string
	NameEn       string
	City         string
	Category     string
	Description  string
	SubmittedBy  string
	DashboardURL string
}

// NotifyNewUniversity sends a beautiful notification for new university
func NotifyNewUniversity(data entity.PendingUniversity) error {

	university := "🎓"
	city := "🏙️"
	category := "🏛️"
	user := "👤"
	desc := "📝"
	link := "🔗"
	check := "✅"

	message := fmt.Sprintf(`%s *دانشگاه جدید برای تأیید* %s

%s *نام دانشگاه:*
`+"`"+`%s`+"`"+`

%s *نام انگلیسی:*
`+"`"+`%s`+"`"+`

%s *شهر:*
`+"`"+`%s`+"`"+`

%s *دسته‌بندی:*
`+"`"+`%s`+"`"+`

%s *ارسال کننده:*
`+"`"+`%s`+"`"+`

%s *توضیحات:*
`+"`"+`%s`+"`"+`

%s [مشاهده در پنل مدیریت](%s)

%s لطفاً دانشگاه را بررسی و تأیید کنید.
`,
		university, university,
		university, data.Name,
		university, data.NameEnglish,
		city, data.City,
		category, data.Category,
		user, data.SubmittedBy,
		desc, data.Description,
		link, data.DescriptionEnglish,
		check,
	)
	return Notify(message)
}
