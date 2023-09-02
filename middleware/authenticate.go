package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/blacheinc/pixel/helper"
	userRepository "github.com/blacheinc/pixel/repository/v1/user"
	"github.com/blacheinc/pixel/types"
	"github.com/opensaucerer/barf"
)

func Auth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		barf.Logger().Debug("tough times")

		// get auth header
		authHeader := r.Header.Get("Authorization")

		// validate auth header
		if authHeader != "" {

			// split auth header
			authValue := strings.Split(authHeader, "Bearer ")

			// validate auth header split
			if len(authValue) == 2 {

				// get token from auth header
				token := authValue[1]

				// validate token
				if token != "" {

					// validate jwt token
					claim, valid := helper.VerifyJWT(token)
					if valid {

						user := userRepository.User{}

						// get user from token
						if err := user.FByKeyVal("id", claim.ID); err == nil {

							// set user in context
							r = r.WithContext(context.WithValue(r.Context(), types.AuthCtxKey{}, &user))

							next.ServeHTTP(w, r)
							return
						}

					}

				}

			}
		}

		barf.Logger().Debug("tougher times")

		barf.Response(w).Status(http.StatusUnauthorized).JSON(barf.Res{
			Status:  false,
			Message: "Please login to continue.",
		})
	})

}
