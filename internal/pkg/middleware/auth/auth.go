package auth

import "context"

type AuthProvider interface {
	Auth(ctx context.Context, token string) (userID string, err error)
}
