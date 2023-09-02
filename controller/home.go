package controller

import (
	"net/http"

	"github.com/blacheinc/pixel/types"
	"github.com/opensaucerer/barf"
)

func Home(w http.ResponseWriter, r *http.Request) {

	data := types.Home{
		Status:      true,
		Version:     "1.0.0",
		Name:        "Black Pixels",
		Description: "The No. 1 stock platform dedicated to promoting African stock photos videos and viusal art.",
		Website:     "https://blackpixels.app",
	}

	barf.Response(w).Status(http.StatusOK).JSON(barf.Res{
		Status:  true,
		Data:    data,
		Message: "Black Pixels",
	})
}
