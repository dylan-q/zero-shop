package middleware

import (
	"fmt"
	"net/http"
)

type LogMiddleware struct {
}

func NewLogMiddleware() *LogMiddleware {
	return &LogMiddleware{}
}

func (m *LogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		fmt.Println("经过log 中间件")
		// Pass through to next handler if need
		next(w, r)
	}
}
