package user

import "context"

type contextKey struct{}

// SetUser to context.
func SetUser(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, contextKey{}, user)
}

// GetUser from context.
func GetUser(ctx context.Context) (*User, bool) {
	user, ok := ctx.Value(contextKey{}).(*User)
	return user, ok
}
