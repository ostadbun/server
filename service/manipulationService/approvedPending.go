package manipulationService

import (
	"context"
)

func (m Manipulation) ApprvingLesson(ctx context.Context, pendingLessionId, userId int64) error {

	return m.manipulationRepo.ApproveLesson(ctx, pendingLessionId, userId)

}

func (m Manipulation) RejectLesson(ctx context.Context, reason *string, pendingLessionId, userId int64) error {

	return m.manipulationRepo.RejectLesson(ctx, pendingLessionId, userId, reason)

}

//--------

func (m Manipulation) ApprvingProfessor(ctx context.Context, pendingprofessorId, userId int64) error {

	return m.manipulationRepo.ApproveProfessor(ctx, pendingprofessorId, userId)

}

func (m Manipulation) RejectProfessor(ctx context.Context, reason *string, pendingprofessorId, userId int64) error {

	return m.manipulationRepo.RejectProfessor(ctx, pendingprofessorId, userId, reason)

}

// --------
func (m Manipulation) ApprvingUniversity(ctx context.Context, pendinguniversityId, userId int64) error {

	return m.manipulationRepo.ApproveUniversity(ctx, pendinguniversityId, userId)

}

func (m Manipulation) RejectUniversity(ctx context.Context, reason *string, pendinguniversityId, userId int64) error {

	return m.manipulationRepo.RejectUniversity(ctx, pendinguniversityId, userId, reason)

}

// --------
func (m Manipulation) ApprvingMajor(ctx context.Context, pendingMajorId, userId int64) error {

	return m.manipulationRepo.ApproveMajor(ctx, pendingMajorId, userId)

}

func (m Manipulation) RejectMajor(ctx context.Context, reason *string, pendingMajorId, userId int64) error {

	return m.manipulationRepo.RejectMajor(ctx, pendingMajorId, userId, reason)
}
