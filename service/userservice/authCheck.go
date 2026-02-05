package userservice

import (
	"context"
)

func (r User) AuthCheck(ctx context.Context, token string) (int, error) {

	return r.redis.AuthCheck(ctx, token)

}

func (r User) RemoveState(ctx context.Context, state string) {

	r.redis.RemoveState(ctx, state)

}

func (r User) CheckIntiRedis(ctx context.Context, state string) ([]byte, string, error) {
	return r.redis.CheckIntoRedis(ctx, state)
}
