package homecontroller

import (
	"crud_go_native/helpers"
	"html/template"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/home/index.gohtml")
	helpers.FuncError(err, "ParseFile Error at home controller Welcome")

	temp.Execute(w, nil)
}
