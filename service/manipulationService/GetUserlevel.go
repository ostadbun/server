package manipulationService

func (m Manipulation) GetUserLevel(userId int) (int, error) {

	return m.activity.LevelCalculator(userId)

}
