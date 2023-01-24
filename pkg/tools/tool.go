package tools

import (
	"go-gin-demo/pkg/setting"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	result := 0
	var page int = 0
	var err error
	if arg := c.Query("page"); arg != "" {
		page, err = strconv.Atoi(arg)
		if err != nil {
			page = 0
			log.Println("[ERROR] ", err)
		}
	}
	if page > 0 {
		result = (page - 1) * setting.ENV.PageSize
	}

	return result
}
