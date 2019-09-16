package middleware

import (
	"github.com/gin-gonic/gin"
	"ibofdagent/server/routers/mtool/api"
	"ibofdagent/server/util"
)

func CheckMtoolHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		xrid := ctx.GetHeader("X-request-Id")
		ts := ctx.GetHeader("ts")

		//ToDo: description -> Enum or (Panic and Recovery)
		if util.IsValidUUID(xrid) == false {
			description := "X-request-Id is invalid in header"
			api.ReturnFail(ctx, description, 10240)
		}

		if ts == "" {
			description := "ts is missing in header"
			api.ReturnFail(ctx, description, 10250)
		}

		ctx.Header("Content-Type", "application/json; charset=utf-8")
		ctx.Header("X-request-Id", xrid)

		ctx.Next()
	}
}
