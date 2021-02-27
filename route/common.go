package route

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/postmy/app/delivery"
)

// Common common delivery
func Common(delivery *delivery.Deliveries) {
	web.Router("/", delivery, "get:Index")
	web.Router("/ping", delivery, "get:Ping")
}
