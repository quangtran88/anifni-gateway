package utils

import (
	"context"
	"time"
)

func InitGRPCContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second)
}
