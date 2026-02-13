package notify

import (
	"fmt"
	"ostadbun/entity"
)

// ProfessorNotificationData contains professor information for notification
type ProfessorNotificationData struct {
	Name         string
	NameEn       string
	Description  string
	SubmittedBy  string
	DashboardURL string
}

// NotifyNewProfessor sends a beautiful notification for new professor
func NotifyNewProfessor(data entity.PendingProfessor) error {

	professor := "👨‍🏫"
	user := "👤"
	desc := "📝"
	link := "🔗"
	check := "✅"

	message := fmt.Sprintf(`%s *استاد جدید برای تأیید* %s

%s *نام استاد:*
`+"`"+`%s`+"`"+`

%s *نام انگلیسی:*
`+"`"+`%s`+"`"+`

%s *ارسال کننده:*
`+"`"+`%d`+"`"+`

%s *توضیحات:*
`+"`"+`%s`+"`"+`

%s [مشاهده در پنل مدیریت](%s)

![%s](%s)

%s لطفاً استاد را بررسی و تأیید کنید.
`,
		professor, professor,
		professor, data.Name,
		professor, data.NameEnglish,
		user, data.SubmittedBy,
		desc, data.Description,
		link, data.DescriptionEnglish,
		data.NameEnglish, *data.ImageUrl,
		check,
	)

	return Notify(message)
}
