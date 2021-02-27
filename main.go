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

	fmt.Printf("Starting server at port %s\n", os.Getenv("PORT"))
	web.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello!")
	// })

	// fmt.Printf("Starting server at port %s\n", os.Getenv("PORT"))

	// if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil); err != nil {
	// 	log.Fatal(err)
	// }

}
