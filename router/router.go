package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type FnRegistRouter func(rg *gin.RouterGroup) //注册路由函数，rg表示路由的根path。

var gfnRouters []FnRegistRouter

func RegistRouter(fn FnRegistRouter) { //将需要注册的路由函数放入数组中，在后面的InitRouter函数中一起注册。
	if fn == nil {
		return
	}
	gfnRouters = append(gfnRouters, fn)
}

func InitRouter() {
	//初始化gin。
	r := gin.Default()

	//从配置文件中读取port信息和路由根path信息。
	port := ":" + viper.GetString("Server.port")
	rootgroup := r.Group(viper.GetString("Server.rootgroup"))

	fmt.Println(viper.GetString("Server.rootgroup"))

	//初始化基础路由，将需要注册的路由都存入 gfnRouters 数组中。
	InitBaseRouters()
	for _, fn := range gfnRouters {
		fn(rootgroup)
	}
	routers := r.Routes()
	for _, v := range routers {
		fmt.Println(v.Path)
	}

	//GIN，启动！！
	r.Run(port)
}

func InitBaseRouters() {
	InitCoreRouter()
}