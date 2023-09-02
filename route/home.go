package route

import (
	"github.com/blacheinc/pixel/controller"
	"github.com/opensaucerer/barf"
)

func RegisterHomeRoutes(frame *barf.SubRoute) {

	frame.Get("/", controller.Home)
}
