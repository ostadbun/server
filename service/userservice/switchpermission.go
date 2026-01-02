package userservice

func (r User) SwitchPermission(userID, masterId int) (bool, error) {

	return r.repo.SwitchPermission(userID, masterId)

}
