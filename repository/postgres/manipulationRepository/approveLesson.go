package manipulationRepository

import (
	"context"
	"errors"
	"time"
)

var (
	ErrLessonNotFound      = errors.New("lesson not found")
	ErrInvalidLessonStatus = errors.New("invalid status for lesson: must be 'approved' or 'rejected'")
)

// ApproveLesson approves a pending lesson
func (d DB) ApproveLesson(ctx context.Context, pendingLessonID int64, approvedBy int64) error {
	return d.updateLessonStatus(ctx, pendingLessonID, "approved", approvedBy, nil)
}

// RejectLesson rejects a pending lesson with optional rejection reason
func (d DB) RejectLesson(ctx context.Context, pendingLessonID int64, rejectedBy int64, rejectionReason *string) error {
	return d.updateLessonStatus(ctx, pendingLessonID, "rejected", rejectedBy, rejectionReason)
}

// updateLessonStatus is a helper method for updating lesson status
func (d DB) updateLessonStatus(
	ctx context.Context,
	pendingLessonID int64,
	status string,
	approvedBy int64,
	rejectionReason *string,
) error {
	if status != "approved" && status != "rejected" {
		return ErrInvalidLessonStatus
	}

	query := `
		UPDATE pending_lesson 
		SET 
			status = $1,
			approved_by = $2,
			approved_at = $3,
			rejection_reason = $4
		WHERE id = $5 AND status = 'pending'
	`

	result, err := d.conn.Conn().ExecContext(ctx, query,
		status,
		approvedBy,
		time.Now().UTC(),
		rejectionReason,
		pendingLessonID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrLessonNotFound
	}

	return nil
}
