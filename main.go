package main

import (
	"ginnote/routers"
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
)

func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	r := gin.Default()

	//自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})

	//加载模板
	r.LoadHTMLGlob("templates/**/*")
	//这里括号里参数的意思就是加载目录"templates/xxx/"下的所有模板。
	//如果换成"templates/*"就变成的加载目录"templates/"下的所有模板。

	//配置静态web服务，第一个参数表示路由，第二个参数表示映射的目录
	r.Static("/static", "./static")

	routers.AdminRoutersInit(r) //调用路由组
	routers.DefaultRoutersInit(r)
	routers.ApiRoutersInit(r)
	r.Run() //默认端口是8000。括号里可以指定端口号，接收字符串类型，比如":4000"。注意":"不能省略！
}
