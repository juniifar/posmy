package route

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/postmy/app/delivery"
)

//Upload xxxx
func Upload(delivery *delivery.Deliveries) {
	web.Router("/v1/upload", delivery, "post:Upload")
	web.Router("/v1/download", delivery, "post:Download")
}
