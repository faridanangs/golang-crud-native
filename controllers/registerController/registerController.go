package registercontroller

import (
	"crud_go_native/entities"
	"crud_go_native/helpers"
	registermodel "crud_go_native/models/registerModel"
	"html/template"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/register/signUp.htm")
		helpers.FuncError(err, "Error parse file at sign up controller get")
		temp.Execute(w, nil)
	}
	if r.Method == "POST" {
		var user entities.Register

		user.Name = r.FormValue("name")
		user.Password = r.FormValue("password")

		if ok := registermodel.Create(user); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/register/login", http.StatusSeeOther)
	}
}

// func LogIn(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		temp, err := template.ParseFiles("views/register/logIn.htm")
// 		helpers.FuncError(err, "Error parse file at log in controller get")
// 		temp.Execute(w, nil)
// 	} else if r.Method == "POST" {
// 		var user entities.Register
// 		users := registermodel.UserAll()

// 		user.Name = r.FormValue("name")
// 		user.Password = r.FormValue("password")

// 		for _, data := range users {
// 			if data.Name == user.Name && data.Password == user.Password {
// 				http.Redirect(w, r, "/products", http.StatusSeeOther)
// 				return
// 			}
// 		}

// 		temp, err := template.ParseFiles("views/register/logIn.htm")
// 		helpers.FuncError(err, "Error parse file at log in controller get")

// 		temp.Execute(w, map[string]any{
// 			"message": "Data tidak di temukan",
// 		})

// 		token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
// 			"name": user.Name,
// 			"exp":  time.Now().Add(60 * time.Minute).Unix(),
// 		})

// 		t, err := token.SignedString([]byte("rahasia"))

// 		if _, err := r.Cookie("token"); err == http.ErrNoCookie {
// 			cookie := &http.Cookie{
// 				Name:     "token",
// 				Value:    t,
// 				Expires:  time.Now().Add(30 * time.Minute),
// 				HttpOnly: true,
// 			}
// 			http.SetCookie(w, cookie)
// 		}

// 	}
// }

// ... (kode sebelumnya)

func LogIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/register/logIn.htm")
		helpers.FuncError(err, "Error parse file at log in controller get")
		temp.Execute(w, nil)
	} else if r.Method == "POST" {
		var user entities.Register
		users := registermodel.UserAll()

		user.Name = r.FormValue("name")
		user.Password = r.FormValue("password")

		for _, data := range users {
			if data.Name == user.Name && data.Password == user.Password {
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"name": user.Name,
					"exp":  time.Now().Add(60 * time.Minute).Unix(),
				})

				t, err := token.SignedString([]byte("rahasia"))
				if err != nil {
					// Tangani kesalahan pembuatan token
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}

				if _, err := r.Cookie("token"); err == http.ErrNoCookie {
					cookie := &http.Cookie{
						Name:     "token",
						Value:    t,
						Expires:  time.Now().Add(30 * time.Minute),
						HttpOnly: true,
						Path: "/",
					}
					http.SetCookie(w, cookie)
				}

				http.Redirect(w, r, "/products", http.StatusSeeOther)
				return
			}
		}

		temp, err := template.ParseFiles("views/register/logIn.htm")
		helpers.FuncError(err, "Error parse file at log in controller get")

		temp.Execute(w, map[string]interface{}{
			"message": "Data tidak di temukan",
		})
	}
}

// ... (kode setelahnya)
