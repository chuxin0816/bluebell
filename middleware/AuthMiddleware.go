package middleware

import (
	"context"
	"strings"

	"github.com/chuxin0816/Scaffold/dao/mysql"
	"github.com/chuxin0816/Scaffold/pkg/jwt"
	"github.com/chuxin0816/Scaffold/response"
	"github.com/cloudwego/hertz/pkg/app"
)

func AuthMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		tokenString := string(ctx.GetHeader("Authorization"))
		// 判断token是否为空或者格式是否正确
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			response.Error(ctx, response.CodeNoAuthority, "")
			ctx.Abort()
			return
		}
		// 去掉token前缀
		tokenString = tokenString[7:]
		// 解析token
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			response.Error(ctx, response.CodeNoAuthority, "")
			ctx.Abort()
			return
		}
		userID := claims.UserID
		user, exist := mysql.CheckUserIDExist(userID)
		if !exist {
			response.Error(ctx, response.CodeUserNotExist, "")
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next(c)
	}
}
