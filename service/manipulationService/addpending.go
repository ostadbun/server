package manipulationService

import (
	"ostadbun/entity"
)

func (m Manipulation) AddPendingUniversity(lesson entity.PendingUniversity, userId int) error {

	return m.manipulationRepo.AddUniversityPending(lesson, userId)

}

func (m Manipulation) AddPendingProfessor(lesson entity.PendingProfessor, userId int) error {

	return m.manipulationRepo.AddProfessorPending(lesson, userId)

}

func (m Manipulation) AddPendingLesson(lesson entity.PendingLesson, userId int) error {

	return m.manipulationRepo.AddLessonPending(lesson, userId)

}

func (m Manipulation) AddPendingMajor(lesson entity.PendingMajor, userId int) error {

	return m.manipulationRepo.AddMajorPending(lesson, userId)

}
