package router

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

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

	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	//初始化gin。
	r := gin.Default()

	//从配置文件中读取port信息和路由根path信息。
	port := ":" + viper.GetString("Server.port")
	if port == ":" { //若配置文件没读取到
		port += "9999"
	}
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

	server := &http.Server{
		Addr:    port,
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Start Server Error: ", err.Error())
			return
		}
	}()

	<-ctx.Done()

	ctx, cancelShutDown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutDown()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown Error: ", err.Error())
	}
	// //GIN，启动！！
	// err := r.Run(port)
	// if err != nil {
	// 	panic(fmt.Sprintf("Start Server Error: %v", err.Error()))
	// }
	fmt.Println("Server ShutDown Success!")
}

func InitBaseRouters() {
	InitCoreRouter()
}
