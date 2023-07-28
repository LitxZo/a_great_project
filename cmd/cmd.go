package cmd

import (
	"fmt"
	"greatproject/conf"
	"greatproject/router"
)

func Start() {
	conf.InitConfig()
	router.InitRouter()
}

func Close() {
	fmt.Println("=================== Clean ==================")
}
