package notify

import (
	"fmt"
	"ostadbun/entity"
)

// MajorNotificationData contains major information for notification
type MajorNotificationData struct {
	Name         string
	NameEn       string
	Description  string
	SubmittedBy  string
	DashboardURL string
}

// NotifyNewMajor sends a beautiful notification for new major
func NotifyNewMajor(data entity.PendingMajor) error {

	major := "🎓"
	book := "📚"
	user := "👤"
	desc := "📝"
	link := "🔗"
	check := "✅"

	message := fmt.Sprintf(`%s *رشته تحصیلی جدید برای تأیید* %s

%s *نام رشته:*
`+"`"+`%s`+"`"+`

%s *نام انگلیسی:*
`+"`"+`%s`+"`"+`

%s *ارسال کننده:*
`+"`"+`%d`+"`"+`

%s *توضیحات:*
`+"`"+`%s`+"`"+`

%s [مشاهده در پنل مدیریت](%s)

%s لطفاً رشته را بررسی و تأیید کنید.
`,
		book, book,
		major, data.Name,
		major, data.NameEnglish,
		user, data.SubmittedBy,
		desc, data.Description,
		link, data.DescriptionEnglish,
		check,
	)

	return Notify(message)
}
