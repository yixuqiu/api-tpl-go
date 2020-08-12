package middlewares

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/shenghui0779/yiigo"
	"go.uber.org/zap"
)

// Recovery panic recover middleware
func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			// panic 捕获
			if err := recover(); err != nil {
				yiigo.Logger().Error(fmt.Sprintf("pay-center panic: %v", err),
					zap.String("request_id", ctx.GetHeader("request_id")),
					zap.String("stack", string(debug.Stack())),
				)

				ctx.JSON(http.StatusOK, gin.H{
					"success": false,
					"code":    50000,
					"msg":     "服务器错误，请稍后重试",
				})

				ctx.Abort()

				return
			}
		}()

		ctx.Next()
	}
}
