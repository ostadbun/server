package academic

import (
	"ostadbun/entity"

	"github.com/gofiber/fiber/v2"
)

type academic struct {
	University []entity.University `json:"university,omitempty"`
	Lesson     []entity.Lesson     `json:"lesson,omitempty"`
	Professor  []entity.Professor  `json:"professor,omitempty"`
	Major      []entity.Major      `json:"major,omitempty"`
}

func (h Handler) Academics(c *fiber.Ctx) error {
	university := c.Query("university")
	lesson := c.Query("lesson")
	professor := c.Query("professor")
	major := c.Query("major")

	var academicsData academic
	var errU, errL, errM, errP error

	var Universities []entity.University
	var Lessons []entity.Lesson
	var Professors []entity.Professor
	var Majors []entity.Major

	if len(university) > 0 {
		Universities, errU = h.academicService.UniversitySearch(university)
	}

	if len(lesson) > 0 {
		Lessons, errL = h.academicService.LessonSearch(lesson)
	}

	if len(professor) > 0 {
		Professors, errP = h.academicService.ProfessorSearch(professor)
	}

	if len(major) > 0 {
		Majors, errM = h.academicService.MajorSearch(major)
	}

	if Universities != nil && errU == nil {
		academicsData.University = Universities
	}

	if Lessons != nil && errL == nil {
		academicsData.Lesson = Lessons
	}

	if Majors != nil && errM == nil {

		academicsData.Major = Majors
	}

	if Professors != nil && errP == nil {

		academicsData.Professor = Professors
	}

	return c.Status(200).JSON(academicsData)

}
