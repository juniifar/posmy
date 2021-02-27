package delivery

import "github.com/beego/beego/v2/server/web"

// Deliveries is delivery dependencies
type Deliveries struct {
	web.Controller
}

// New is delivery constructor
func New() *Deliveries {
	return &Deliveries{}
}
