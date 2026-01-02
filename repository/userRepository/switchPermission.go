package userRepository

import "database/sql"

func (a DB) SwitchPermission(userID, masterID int) (bool, error) {

	var adminBy sql.NullInt64

	query := `UPDATE users 
              SET admin_by = CASE WHEN admin_by IS NULL THEN $2::int ELSE NULL END 
              WHERE id = $1 
              RETURNING admin_by`

	err := a.conn.Conn().QueryRow(query, userID, masterID).Scan(&adminBy)

	if err != nil {
		return false, err
	}

	return adminBy.Valid, nil
}
