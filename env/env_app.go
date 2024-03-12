package env

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
)

type ctxKey struct{}

type AppEnv struct {
	env map[string]any
}

func (app AppEnv) Get(key string) any {
	return app.env[key]
}

func GetAppEnv(key string) string {
	return os.Getenv(key)
}

func GetAppContext() context.Context {
	appCtx := context.Background()
	appCtx = context.WithValue(appCtx, ctxKey{}, AppEnv{
		env: map[string]any{},
	})
	return appCtx
}

func GetAppKey(appCtx context.Context, key string) any {
	return appCtx.Value(ctxKey{}).(AppEnv).env[key]
}

func SetAppKey(appCtx context.Context, key string, value any) {
	appCtx.Value(ctxKey{}).(AppEnv).env[key] = value
}

func SetContext(appCtx context.Context, webEngine *gin.Engine) {
	webEngine.Use(func(ctx *gin.Context) {
		ctx.Set("app_ctx", appCtx)
	})
}

func GetContext(ctx *gin.Context) context.Context {
	appCtx, exist := ctx.Get("app_ctx")
	if exist {
		return appCtx.(context.Context)
	} else {
		return nil
	}
}
