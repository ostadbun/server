package academicservice

import (
	"ostadbun/entity"
	"ostadbun/repository/postgres/academicRepository"
)

type Service struct {
	academicRepo academicRepository.DB
}

func New(academicRepo academicRepository.DB) Service {
	return Service{academicRepo: academicRepo}
}

func (s Service) UniversitySearch(name string) ([]entity.University, error) {

	data, err := s.academicRepo.UniversitySearch(name)

	if err != nil {
		// TODO log here
	}

	return data, err

}

func (s Service) ProfessorSearch(name string) ([]entity.Professor, error) {

	data, err := s.academicRepo.ProfessorSearch(name)

	if err != nil {
		// TODO log here
	}

	return data, err

}

func (s Service) LessonSearch(name string) ([]entity.Lesson, error) {

	data, err := s.academicRepo.LessonSearch(name)

	if err != nil {
		// TODO log here
	}

	return data, err

}

func (s Service) MajorSearch(name string) ([]entity.Major, error) {

	data, err := s.academicRepo.MajorSearch(name)

	if err != nil {
		// TODO log here
	}

	return data, err

}
