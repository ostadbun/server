package userservice

func (r User) AdminChecker(userID string) error {

	_, err := r.repo.AdminByWho(userID)

	return err
}
