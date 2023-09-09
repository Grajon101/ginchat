package router

import (
	"ginchat/docs"
	"ginchat/service"
	"html/template"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// gin-swagger middleware
// swagger embed files

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/index", service.GetIndex)
	userGroup := r.Group("/user")
	{
		userGroup.GET("/getUserList", service.GetUserList)
		userGroup.GET("/createUser", service.CreateUser)
		userGroup.GET("/deleteUser", service.DeleteUser)
		userGroup.POST("/updateUser", service.UpdateUser)
		userGroup.GET("/getUserByNameAndPwd", service.GetUserByNameAndPwd)
	}
	//发送消息
	r.GET("/user/sendMsg", service.SendMsg)

	r.POST("/postPerson", service.PostPerson)

	r.GET("/a", func(c *gin.Context) {

		ind, err := template.ParseFiles("web/a.html",
			"web/b.html")
		if err != nil {
			panic(err)
		}
		ind.Execute(c.Writer, "user")

	})

	return r
}
