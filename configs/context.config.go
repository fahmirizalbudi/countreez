package configs

import "context"

var Ctx context.Context

func GetCtx() {
	Ctx = context.Background()
}