package main

import (
	"fmt"
	"os"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/postmy/app/delivery"
	"github.com/postmy/route"
)

func main() {

	delivery := delivery.New()
	logs.Info("initialize delivery")

	route.Common(delivery)
	route.Upload(delivery)

	fmt.Printf("Starting server at port %s\n", os.Getenv("PORT"))
	web.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))

}
