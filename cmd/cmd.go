package cmd

import (
	"fmt"
	"greatproject/conf"
)

func Start() {
	conf.InitConfig()
}

func Close() {
	fmt.Println("=================== Clean ==================")
}
