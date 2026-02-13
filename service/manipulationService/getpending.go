package manipulationService

import (
	"ostadbun/entity"
)

func (m Manipulation) GetPendingUniversity() ([]entity.PendingUniversity, error) {

	return m.manipulationRepo.GetUniversityPending()

}

func (m Manipulation) GetPendingProfessor() ([]entity.PendingProfessor, error) {

	return m.manipulationRepo.GetProfessorPending()

}

func (m Manipulation) GetPendingLesson() ([]entity.PendingLesson, error) {

	return m.manipulationRepo.GetLessonPending()

}

func (m Manipulation) GetPendingMajor() ([]entity.PendingMajor, error) {

	return m.manipulationRepo.GetMajorPending()

}
