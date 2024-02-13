package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
}

func (con ArticleController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "admin/news.html", gin.H{})
}

func (con ArticleController) Add(ctx *gin.Context) {
	ctx.String(http.StatusOK, "增加文章--")
}

func (con ArticleController) Edit(ctx *gin.Context) {
	ctx.String(http.StatusOK, "修改文章--")
}
