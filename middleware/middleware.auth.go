package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/wafellofazztrack/lectronic-backend/lib"
)

func AuthMiddleware(role ...string) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var header string
			var valid bool

			// get header optional
			if header = r.Header.Get("Authorization"); header == "" {
				lib.NewRes("header not provided, please login", 401, true).Send(w)
				return
			}

			if !strings.Contains(header, "Bearer") {
				lib.NewRes("invalid header type", 401, true).Send(w)
				return
			}

			tokens := strings.Replace(header, "Bearer ", "", -1)

			checkToken, err := lib.CheckToken(tokens)
			if err != nil {
				lib.NewRes(err.Error(), 401, true).Send(w)
				return
			}

			for _, rl := range role {
				if rl == checkToken.Role {
					valid = true
				}
			}

			if !valid {
				lib.NewRes("you do not have permission to access", 401, true).Send(w)
				return
			}

			log.Println("Auth middleware pass")

			// share userid to controller
			ctx := context.WithValue(r.Context(), "user", checkToken.UserID)

			// serve next middleware
			next.ServeHTTP(w, r.WithContext(ctx))

		})
	}
}
