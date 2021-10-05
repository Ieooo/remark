package main

import (
	"fmt"
	"joke/conf"
	"joke/connection"
	"joke/router"
)

func main() {
	// load config.ini
	err := conf.Init("conf/config.ini")
	if err != nil {
		fmt.Printf("error :%v", err)
	}

	// connect database
	connection.Connect(conf.Conf.MysqlConfig)

	// startup router
	r := router.RouterInit()
	if err := r.Run(fmt.Sprintf(":%d", conf.Conf.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
