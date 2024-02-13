package defaults

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultController struct {
}

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (con DefaultController) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "default/index.html", gin.H{
		"title": "首页",
		"msg":   "我是msg",
		"score": 89,
		"hobby": []string{"吃饭", "睡觉", "写代码"},
		"newsList": []interface{}{
			Article{
				Title:   "新闻标题1",
				Content: "新闻内容1",
			},
			Article{
				Title:   "新闻标题2",
				Content: "新闻内容2",
			},
		},
		"testSlice": []string{},
		"news": Article{
			Title:   "新闻标题3",
			Content: "新闻内容3",
		},
		"date": 1707723555,
	})
}

func (con DefaultController) News(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "default/news.html", gin.H{
		"title": "新闻首页",
		"news": Article{
			Title:   "新闻标题",
			Content: "新闻内容",
		},
	})
}

func (con DefaultController) User(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "default/user.html", gin.H{})
}

type UserInfo struct {
	Username string `json:"username" form:"username"` //form是浏览器链接?username=xxxx
	Password string `json:"password" form:"password"`
}

func (con DefaultController) DoAddUser(ctx *gin.Context) {
	user := UserInfo{}

	if err := ctx.ShouldBind(&user); err == nil {
		fmt.Printf("user: %v\n", user)
		ctx.JSON(http.StatusOK, gin.H{
			"username": user.Username, //""里的是返回值，:后的是变量
			"password": user.Password,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
	}
}
