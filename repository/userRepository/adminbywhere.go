package userRepository

func (a AuthRepo) AdminByWho(userID string) (int, error) {

	var userIdFromDB int

	err := a.db.QueryRow("select admin_by from users where id=$1", userID).Scan(&userIdFromDB)

	if err != nil {
		return 0, err
	}

	return userIdFromDB, nil
}
