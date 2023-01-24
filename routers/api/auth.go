package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-gin-demo/models"
	"go-gin-demo/pkg/errorcode"
	"go-gin-demo/pkg/tools"
)

// @Summary Get Auth
// @Produce  json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {string} json "{"code":200,"data":{token:"..."},"msg":"..."}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"error msg"}"
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if len(username) > 50 || len(username) == 0 {
		code := errorcode.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errorcode.GetMsg(code),
		})
		return
	}
	if len(password) > 50 || len(password) == 0 {
		code := errorcode.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errorcode.GetMsg(code),
		})
		return
	}

	data := make(map[string]interface{})
	code := errorcode.SUCCESS

	isExist, _ := models.CheckAuth(username, password) //確認帳密皆有在DB
	if isExist {
		token, err := tools.GenerateToken(username, password) //產生Token
		if err != nil {
			code = errorcode.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token
			code = errorcode.SUCCESS
		}

	} else {
		code = errorcode.ERROR_AUTH
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errorcode.GetMsg(code),
		"data": data,
	})
}
