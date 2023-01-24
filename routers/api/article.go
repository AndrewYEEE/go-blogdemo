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

//获取单个文章
func GetArticle(c *gin.Context) {
	log.Println("[INFO] GetArticle API")
	var err error
	var id int
	var code int
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("[ERROR] Get id error!!!", err)
		code = errorcode.ERROR_ADD_ARTICLE_FAIL
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
	var data interface{}
	exist, _ := models.CheckExistArticleByID(id)
	if exist {
		data, err = models.GetArticleById(id)
		if err != nil {
			code = errorcode.ERROR_NOT_EXIST_ARTICLE
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errorcode.GetMsg(code),
			})
			return
		}
		code = errorcode.SUCCESS
	} else {
		code = errorcode.ERROR_NOT_EXIST_ARTICLE
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errorcode.GetMsg(code),
		"data": data,
	})
}

//获取多个文章
func GetArticles(c *gin.Context) {
	log.Println("[INFO] GetArticles API")
	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	var state int = -1
	var err error
	if arg := c.Query("state"); arg != "" {
		state, err = strconv.Atoi(arg)
		if err != nil {
			log.Println("[ERROR] state", err)
			code := errorcode.ERROR
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errorcode.GetMsg(code),
			})
			return
		}
		maps["state"] = state
	}

	if state < 0 && state > 1 {
		log.Println("[ERROR] state只允许0或1")
		code := errorcode.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errorcode.GetMsg(code),
		})
		return
	}

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId, err = strconv.Atoi(arg)
		if err != nil {
			log.Println("[ERROR] tag_id", err)
			code := errorcode.ERROR
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errorcode.GetMsg(code),
			})
			return
		}
		maps["tag_id"] = tagId
	}

	if tagId <= 0 {
		log.Println("[ERROR] tagId必须大于0")
		code := errorcode.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errorcode.GetMsg(code),
		})
		return
	}

	code := errorcode.SUCCESS

	data["lists"], err = models.GetArticlesByPages(tools.GetPage(c), setting.ENV.PageSize, maps)
	if err != nil {
		log.Println("[ERROR] GetArticlesByPages ", err)
		code := errorcode.ERROR_GET_ARTICLES_FAIL
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errorcode.GetMsg(code),
		})
		return
	}
	data["total"], err = models.GetTotalArticle(maps)
	if err != nil {
		log.Println("[ERROR] GetTotalArticle ", err)
		code := errorcode.ERROR_GET_ARTICLES_FAIL
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errorcode.GetMsg(code),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errorcode.GetMsg(code),
		"data": data,
	})
}

//新增文章
func AddArticle(c *gin.Context) {
	log.Println("[INFO] AddArticle API")
	var tagId int = -1
	var err error
	if arg := c.Query("tag_id"); arg != "" {
		tagId, err = strconv.Atoi(arg)
		if err != nil {
			log.Println("[ERROR] tag_id", err)
			code := errorcode.ERROR
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errorcode.GetMsg(code),
			})
			return
		}
	}
	if tagId <= 0 {
		log.Println("[ERROR] ID必须大于0")
		code := errorcode.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errorcode.GetMsg(code),
		})
		return
	}

	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	var state int
	state, err = strconv.Atoi(c.DefaultQuery("state", "0"))
	if err != nil {
		state = 0
		err = nil
	}

	code := errorcode.SUCCESS
	exist, _ := models.CheckExistTagByID(tagId)
	if exist {
		data := make(map[string]interface{})
		data["tag_id"] = tagId
		data["title"] = title
		data["desc"] = desc
		data["content"] = content
		data["created_by"] = createdBy
		data["state"] = state

		err = models.AddArticle(data)
		if err != nil {
			log.Println("[ERROR] tag_id", err)
			code := errorcode.ERROR_ADD_ARTICLE_FAIL
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
		"data": make(map[string]interface{}),
	})
}

//修改文章
func EditArticle(c *gin.Context) {
	log.Println("[INFO] EditArticle API")
	var err error

	var id int = -1
	if arg := c.Query("id"); arg != "" {
		id, err = strconv.Atoi(arg)
		if err != nil {
			log.Println("[ERROR] id", err)
			code := errorcode.ERROR
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errorcode.GetMsg(code),
			})
			return
		}
	}
	if id <= 0 {
		log.Println("[ERROR] id必须大于0")
		code := errorcode.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errorcode.GetMsg(code),
		})
		return
	}

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId, err = strconv.Atoi(arg)
		if err != nil {
			log.Println("[ERROR] tag_id", err)
			code := errorcode.ERROR
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errorcode.GetMsg(code),
			})
			return
		}
	}
	if tagId <= 0 {
		log.Println("[ERROR] tagId必须大于0")
		code := errorcode.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errorcode.GetMsg(code),
		})
		return
	}

	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state, err = strconv.Atoi(arg)
		if err != nil {
			log.Println("[ERROR] state", err)
			code := errorcode.ERROR
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errorcode.GetMsg(code),
			})
			return
		}
	}
	if state < 0 && state > 1 {
		log.Println("[ERROR] state只允许0或1")
		code := errorcode.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errorcode.GetMsg(code),
		})
		return
	}

	code := errorcode.SUCCESS
	exist, _ := models.CheckExistArticleByID(id)
	if exist {
		tagexist, _ := models.CheckExistTagByID(tagId)
		if tagexist {
			data := make(map[string]interface{})
			if tagId > 0 {
				data["tag_id"] = tagId
			}
			if title != "" {
				data["title"] = title
			}
			if desc != "" {
				data["desc"] = desc
			}
			if content != "" {
				data["content"] = content
			}

			data["modified_by"] = modifiedBy

			err = models.EditArticleById(id, data)
			if err != nil {
				log.Println("[ERROR] EditArticleById ", err)
				code = errorcode.ERROR_EDIT_ARTICLE_FAIL
				c.JSON(http.StatusOK, gin.H{
					"code": code,
					"msg":  errorcode.GetMsg(code),
				})
				return
			}
			code = errorcode.SUCCESS
		} else {
			code = errorcode.ERROR_NOT_EXIST_TAG
		}
	} else {
		code = errorcode.ERROR_NOT_EXIST_ARTICLE
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errorcode.GetMsg(code),
		"data": make(map[string]string),
	})
}

//删除文章
func DeleteArticle(c *gin.Context) {
	log.Println("[INFO] DeleteArticle API")
	var err error
	var id int = -1
	if arg := c.Query("id"); arg != "" {
		id, err = strconv.Atoi(arg)
		if err != nil {
			log.Println("[ERROR] id", err)
			code := errorcode.ERROR
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errorcode.GetMsg(code),
			})
			return
		}
	}
	if id <= 0 {
		log.Println("[ERROR] id必须大于0")
		code := errorcode.INVALID_PARAMS
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errorcode.GetMsg(code),
		})
		return
	}
	code := errorcode.SUCCESS
	exist, _ := models.CheckExistArticleByID(id)
	if exist {
		err = models.DeleteArticleById(id)
		if err != nil {
			log.Println("[ERROR] DeleteArticleById", err)
			code := errorcode.ERROR_DELETE_ARTICLE_FAIL
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errorcode.GetMsg(code),
			})
			return
		}
		code = errorcode.SUCCESS
	} else {
		code = errorcode.ERROR_NOT_EXIST_ARTICLE
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errorcode.GetMsg(code),
		"data": make(map[string]string),
	})
}
