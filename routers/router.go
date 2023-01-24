package routers

import (
	"go-gin-demo/middleware"
	"go-gin-demo/pkg/setting"

	"go-gin-demo/routers/api"

	_ "go-gin-demo/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ENV.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	r.GET("/auth", api.GetAuth) //Token

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT()) //帶入自訂middle，強制檢查Token
	{
		//获取标签列表
		apiv1.GET("/tags", api.GetTags)
		//新建标签
		apiv1.POST("/tags", api.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", api.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", api.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", api.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", api.GetArticle)
		//新建文章
		apiv1.POST("/articles", api.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", api.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", api.DeleteArticle)
	}

	return r
}
