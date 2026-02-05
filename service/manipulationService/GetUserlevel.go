package manipulationService

import "context"

func (m Manipulation) GetUserLevel(ctx context.Context, userId int) (int, error) {

	return m.activity.LevelCalculator(ctx, userId)

}
