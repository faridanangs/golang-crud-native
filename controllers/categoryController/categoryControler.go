package categorycontroller

import (
	"crud_go_native/entities"
	"crud_go_native/helpers"
	categorymodel "crud_go_native/models/categoryModel"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/category/index.htm")
	helpers.FuncError(err, "error template parseFile at Index category controller")

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/create.htm")
		helpers.FuncError(err, "error template parseFile at Add category controller GET")
		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var creteCategory entities.Category

		creteCategory.Name = r.FormValue("name")
		creteCategory.CreatedAt = time.Now()
		creteCategory.UpdatedAt = time.Now()

		if ok := categorymodel.Create(creteCategory); !ok {
			temp, err := template.ParseFiles("views/category/create.htm")
			helpers.FuncError(err, "error template parseFile at Add category controller POST")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}

}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		categoryID := r.URL.Query().Get("id")
		id, err := strconv.Atoi(categoryID)
		helpers.FuncError(err, "error strconv at edit category controller GET")

		temp, err := template.ParseFiles("views/category/edit.htm")
		helpers.FuncError(err, "error template parseFile at edit category controller GET")

		categoryDetail := categorymodel.Detail(id)

		data := map[string]any{
			"category": categoryDetail,
		}

		temp.Execute(w, data)

	}
	if r.Method == "POST" {
		var category entities.Category

		categoryID := r.FormValue("id")
		id, err := strconv.Atoi(categoryID)
		helpers.FuncError(err, "error strconv at update category controller Post")

		category.Name = r.FormValue("name")
		category.UpdatedAt = time.Now()

		if ok := categorymodel.Update(id, category); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	categoryID := r.FormValue("id")
	id, err := strconv.Atoi(categoryID)
	helpers.FuncError(err, "error strconv at delete category controller")

	categorymodel.Delete(id)

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}
