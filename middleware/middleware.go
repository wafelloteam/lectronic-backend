package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Handle(hd http.HandlerFunc, md ...Middleware) http.HandlerFunc {
	for i := len(md); i > 0; i-- {
		hd = md[i-1](hd)
	}
	return hd

}
