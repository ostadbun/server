package manipulationRepository

import (
	"context"
	"errors"
	"time"
)

var (
	ErrProfessorNotFound      = errors.New("professor not found")
	ErrInvalidProfessorStatus = errors.New("invalid status for professor: must be 'approved' or 'rejected'")
)

// ApproveProfessor approves a pending professor
func (d DB) ApproveProfessor(ctx context.Context, pendingProfessorID int64, approvedBy int64) error {
	return d.updateProfessorStatus(ctx, pendingProfessorID, "approved", approvedBy, nil)
}

// RejectProfessor rejects a pending professor with optional rejection reason
func (d DB) RejectProfessor(ctx context.Context, pendingProfessorID int64, rejectedBy int64, rejectionReason *string) error {
	return d.updateProfessorStatus(ctx, pendingProfessorID, "rejected", rejectedBy, rejectionReason)
}

// updateProfessorStatus is a helper method for updating professor status
func (d DB) updateProfessorStatus(
	ctx context.Context,
	pendingProfessorID int64,
	status string,
	approvedBy int64,
	rejectionReason *string,
) error {
	if status != "approved" && status != "rejected" {
		return ErrInvalidProfessorStatus
	}

	query := `
		UPDATE pending_professor 
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
		pendingProfessorID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrProfessorNotFound
	}

	return nil
}
