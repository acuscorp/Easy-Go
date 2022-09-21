package contexts
import (
	"context"
  "fmt"
)

type userKey int

const key userKey = 1

func ContextWithUser(ctx context.Context, user string) context.Context {
	return context.WithValue(ctx, key, user)
}

func UserFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(key).(string)
	return user, ok
}

func InitExtratFromContext() {
  parentCtx := context.Background()
  ctx := context.WithValue(parentCtx, key, "noe")
  user, ok := UserFromContext(ctx)
  if !ok {
    fmt.Println("User not found")
  }

  fmt.Printf("The user is: %s", user)
}