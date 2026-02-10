package academic

import "ostadbun/service/academicservice"

type Handler struct {
	academicService academicservice.Service
}

func New(academicService academicservice.Service) Handler {
	return Handler{
		academicService: academicService,
	}
}
