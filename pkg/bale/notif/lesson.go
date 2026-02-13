package notify

import (
	"fmt"
	"ostadbun/entity"
)

// NotifyNewLesson sends a beautiful notification to Bale channel/group
func NotifyNewLesson(data entity.PendingLesson) error {
	// ساخت متن زیبا با فرمت‌بندی
	// از کاراکترهای خاص برای جداسازی و زیباسازی استفاده می‌کنیم
	star := "⭐"
	book := "📚"
	level := "🎯"
	user := "👤"
	desc := "📝"
	link := "🔗"
	check := "✅"

	// نمایش سطح سختی به صورت ستاره
	difficultyStars := ""
	for i := 0; i < data.Difficulty; i++ {
		difficultyStars += "⭐"
	}
	for i := data.Difficulty; i < 5; i++ {
		difficultyStars += "☆"
	}

	message := fmt.Sprintf(`%s *درس جدید برای تأیید* %s

%s *نام درس:*
`+"`"+`%s`+"`"+`

%s *نام انگلیسی:*
`+"`"+`%s`+"`"+`

%s *سطح سختی:*
`+"`"+`%s`+"`"+` (%d/5)

%s *ارسال کننده:*
`+"`"+`%d`+"`"+`

%s *توضیحات:*
`+"`"+`%s`+"`"+`

%s *توضیحات انگلیسی:*
`+"`"+`%s`+"`"+`


%s [مشاهده در پنل مدیریت](%s)

%s لطفاً درس را بررسی و تأیید کنید.
`,
		star, star,
		book, data.Name,
		book, data.NameEnglish,
		level, difficultyStars, data.Difficulty,
		user, data.SubmittedBy,
		desc, data.Description,
		desc, data.DescriptionEnglish,
		link, "https://ostadbun.tech",
		check,
	)

	return Notify(message)

}
