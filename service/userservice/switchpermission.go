package userservice

func (r User) SwitchPermission(userID, masterId int) error {

	_, err := r.repo.SwitchPermission(userID, masterId)

	return err
}
