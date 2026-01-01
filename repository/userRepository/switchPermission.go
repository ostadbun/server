package userRepository

func (a DB) SwitchPermission(userID, masterID int) (int, error) {
	var userIdFromDB int

	query := `UPDATE users 
              SET admin_by = CASE WHEN admin_by IS NULL THEN $2::int ELSE NULL END 
              WHERE id = $1 
              RETURNING id` // RETURNING id لازم است چون Scan انجام می‌دهید

	err := a.conn.Conn().QueryRow(query, userID, masterID).Scan(&userIdFromDB)

	if err != nil {
		return 0, err
	}

	return userIdFromDB, nil
}
