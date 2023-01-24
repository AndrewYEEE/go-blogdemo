package api

import (
	"go-gin-demo/models"
	"go-gin-demo/pkg/errorcode"
	"go-gin-demo/pkg/setting"
	"go-gin-demo/pkg/tools"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get multiple article tags
// @Produce  json
// @Param name query string false "Name"
// @Param state query int false "State"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"..."}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"error msg"}"
// @Router /api/v1/tags [get]
func GetTags(c *gin.Context) {
	log.Println("[INFO] GetTags API")
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	var err error
	if arg := c.Query("state"); arg != "" {
		state, err = strconv.Atoi(arg)
		if err != nil {
			log.Println("[ERROR] ", err)
			code := errorcode.ERROR
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errorcode.GetMsg(code),
			})
			return
		}
		maps["state"] = state

	}

	code := errorcode.SUCCESS

	data["lists"], _ = models.GetTags(tools.GetPage(c), setting.ENV.PageSize, maps)
	data["total"], _ = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errorcode.GetMsg(code),
		"data": data,
	})
}

// @Summary Add article tag
// @Produce  json
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param created_by body int false "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"..."}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"error msg"}"
// @Router /api/v1/tags [post]
func AddTag(c *gin.Context) {
	log.Println("[INFO] AddTag API")
	var err error
	var state int
	name := c.Query("name")
	state, err = strconv.Atoi(c.DefaultQuery("state", "0"))
	if err != nil {
		state = 0
		err = nil
	}
	createdBy := c.Query("created_by")

	var code int
	var exist bool = false
	exist, _ = models.CheckExistTagByName(name)
	if !exist {
		log.Println("[INFO] AddTag API", name, state, createdBy)
		err = models.AddTag(name, state, createdBy)
		if err != nil {
			code = errorcode.ERROR_ADD_TAG_FAIL
		}
		code = errorcode.SUCCESS
	} else {
		code = errorcode.ERROR_EXIST_TAG
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errorcode.GetMsg(code),
		"data": make(map[string]string),
	})

}

// @Summary Update article tag
// @Produce  json
// @Param id path int true "ID"
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param modified_by body string true "ModifiedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"..."}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"error msg"}"
// @Router /api/v1/tags/{id} [put]
func EditTag(c *gin.Context) {
	log.Println("[INFO] EditTag API")
	var err error
	var id int
	var code int
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("[ERROR] Get id error!!!", err)
		code = errorcode.ERROR_EDIT_TAG_FAIL
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errorcode.GetMsg(code),
		})
		return
	}

	if id <= 0 {
		log.Println("[ERROR] ID必须大于0")
		code = errorcode.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errorcode.GetMsg(code),
		})
		return
	}
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state, err = strconv.Atoi(arg)
		if err != nil {
			log.Println("[ERROR] ", err)
			code := errorcode.ERROR
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errorcode.GetMsg(code),
			})
			return
		}
	}

	code = errorcode.SUCCESS
	exist, _ := models.CheckExistTagByID(id)
	if exist {
		data := make(map[string]interface{})
		data["modified_by"] = modifiedBy
		if name != "" {
			data["name"] = name
		}
		if state != -1 {
			data["state"] = state
		}

		err = models.EditTag(id, data)
		if err != nil {
			code := errorcode.ERROR
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errorcode.GetMsg(code),
			})
			return
		}
	} else {
		code = errorcode.ERROR_NOT_EXIST_TAG
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errorcode.GetMsg(code),
		"data": make(map[string]string),
	})
}

// @Summary Delete article tag
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"..."}"
// @Failure 500 {string} json "{"code":200,"data":{},"msg":"error msg"}"
// @Router /api/v1/tags/{id} [delete]
func DeleteTag(c *gin.Context) {
	log.Println("[INFO] DeleteTag API")
	var err error
	var id int
	var code int
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("[ERROR] Get id error!!!", err)
		code = errorcode.ERROR_EDIT_TAG_FAIL
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errorcode.GetMsg(code),
		})
		return
	}

	if id <= 0 {
		log.Println("[ERROR] ID必须大于0")
		code = errorcode.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errorcode.GetMsg(code),
		})
		return
	}

	code = errorcode.SUCCESS
	exist, _ := models.CheckExistTagByID(id)
	if exist {
		err = models.DeleteTag(id)
		if err != nil {
			code := errorcode.ERROR
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errorcode.GetMsg(code),
			})
			return
		}
	} else {
		code = errorcode.ERROR_NOT_EXIST_TAG
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errorcode.GetMsg(code),
		"data": make(map[string]string),
	})
}
