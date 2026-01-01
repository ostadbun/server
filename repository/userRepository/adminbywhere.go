package userRepository

func (a DB) AdminByWho(userID string) (int, error) {

	var userIdFromDB int

	err := a.conn.Conn().QueryRow("select admin_by from users where id=$1", userID).Scan(&userIdFromDB)

	if err != nil {
		return 0, err
	}

	return userIdFromDB, nil
}
