package middleware

import (
	"fmt"
	"net/http"

	"github.com/karimla/webapp/models"
)

type requireUser struct {
	models.UserService
}

func NewRequireUser(us models.UserService) *requireUser {
	return &requireUser{us}
}

func (mw *requireUser) Apply(next http.Handler) http.HandlerFunc {
	return mw.ApplyFn(next.ServeHTTP)
}

func (mw *requireUser) ApplyFn(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("remember_token")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		user, err := mw.ByRemember(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		fmt.Println(user)

		next(w, r)
	}
}