package routers

import (
	"ginnote/controllers/defaults"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", defaults.DefaultController{}.Index)
		defaultRouters.GET("/news", defaults.DefaultController{}.News)
		defaultRouters.GET("/user", defaults.DefaultController{}.User)
		defaultRouters.POST("/doAddUser", defaults.DefaultController{}.DoAddUser)
	}
}
