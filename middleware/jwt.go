package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go-gin-demo/pkg/errorcode"
	"go-gin-demo/pkg/tools"
)

func JWT() gin.HandlerFunc { //一個JWT middleware會回傳gin.HandlerFunc
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = errorcode.SUCCESS
		token := c.Query("token") //擷取Token
		if token == "" {
			code = errorcode.INVALID_PARAMS
		} else {
			claims, err := tools.ParseToken(token) //呼叫ParseWithClaims解碼Token
			if err != nil {
				code = errorcode.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.RegisteredClaims.ExpiresAt.Unix() { //如果超時
				code = errorcode.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != errorcode.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  errorcode.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
