package middleware

import "context"

type contextKey int

const (
	ContextKeyUserID contextKey = iota
)

func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, ContextKeyUserID, userID)
}

func GetUserID(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(ContextKeyUserID).(string)
	return userID, ok
}
