package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (con IndexController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/index.html", gin.H{})
}
