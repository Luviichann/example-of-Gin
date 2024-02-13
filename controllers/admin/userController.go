package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (con UserController) Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "用户列表--")
}

func (con UserController) Add(ctx *gin.Context) {
	ctx.String(http.StatusOK, "增加用户--")
}

func (con UserController) Edit(ctx *gin.Context) {
	ctx.String(http.StatusOK, "修改用户--")
}
