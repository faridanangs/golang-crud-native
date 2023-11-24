package middleware

import "net/http"

type LogMiddleware struct {
	Handler http.Handler
}

func (middle *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, errRedirect := r.Cookie("already_redirected")
	_, errToken := r.Cookie("token")

	if errRedirect == nil {
		
		middle.Handler.ServeHTTP(w, r)
		return
	}
	if errToken == http.ErrNoCookie {
		http.SetCookie(w, &http.Cookie{
			Name:   "already_redirected",
			Value:  "true",
			Path:   "/register/login",
			MaxAge: 60,
		})
		http.Redirect(w, r, "/register/login", http.StatusSeeOther)
		return
	}
	middle.Handler.ServeHTTP(w, r)
}
