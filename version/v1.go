package version

import (
	"github.com/blacheinc/pixel/middleware"
	"github.com/blacheinc/pixel/route"
	"github.com/opensaucerer/barf"

	"github.com/blacheinc/pixel/route/v1/user"
)

// Version1Routes registers all routes for the v1 version
func V1() {
	frame := barf.RetroFrame("/v2")
	// access to the api is only allowed with a valid token
	barf.Hippocampus(frame).Hijack(middleware.Auth)
	route.RegisterHomeRoutes(frame)
	user.RegisterUserRoutes(frame)
}
