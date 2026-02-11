package manipulationService

import manipulationParam "ostadbun/param/manipulation"

func (m Manipulation) AddLesson(lesson manipulationParam.PendingLesson, userId int) error {

	return m.manipulationRepo.AddLesson(lesson, userId)

}
