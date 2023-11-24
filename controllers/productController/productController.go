package productcontroller

import (
	"crud_go_native/entities"
	"crud_go_native/helpers"
	categorymodel "crud_go_native/models/categoryModel"
	productmodel "crud_go_native/models/productModel"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	products := productmodel.GetAll()
	data := map[string]any{
		"products": products,
	}

	temp, err := template.ParseFiles("views/product/index.htm")
	helpers.FuncError(err, "error template parseFile at Index product controller")

	temp.Execute(w, data)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idDetail := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idDetail)

	product := productmodel.Detail(id)

	data := map[string]any{
		"product": product,
	}

	temp, err := template.ParseFiles("views/product/detail.htm")
	helpers.FuncError(err, "error parsefile at product detail controller")

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/product/create.htm")
		helpers.FuncError(err, "error template parseFile at Add product controller")

		categories := categorymodel.GetAll()
		data := map[string]any{
			"categories": categories,
		}
		temp.Execute(w, data)
	}
	if r.Method == "POST" {
		var product entities.Product
		stock, _ := strconv.Atoi(r.FormValue("stock"))
		category_id, _ := strconv.Atoi(r.FormValue("category_id"))

		product.Name = r.FormValue("name")
		product.Category.Id = uint(category_id)
		product.Stock = int64(stock)
		product.Description = r.FormValue("description")
		product.CreatedAt = time.Now()
		product.UpdatedAt = time.Now()

		if ok := productmodel.Create(product); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		}

		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		idProduct := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idProduct)

		product := productmodel.Detail(id)
		categories := categorymodel.GetAll()

		data := map[string]any{
			"product":    product,
			"categories": categories,
		}

		temp, err := template.ParseFiles("views/product/update.htm")
		helpers.FuncError(err, "error parse file at product update controller GET")
		temp.Execute(w, data)
	}

	if r.Method == "POST" {
		var product entities.Product

		idCategory, _ := strconv.Atoi(r.FormValue("category_id"))
		stock, _ := strconv.Atoi(r.FormValue("stock"))
		id, _ := strconv.Atoi(r.FormValue("id"))

		product.Name = r.FormValue("name")
		product.Category.Id = uint(idCategory)
		product.Stock = int64(stock)
		product.Description = r.FormValue("description")
		product.UpdatedAt = time.Now()

		if ok := productmodel.Update(id, product); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/products", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idProduct)

	if ok := productmodel.Delete(id); !ok {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/products", http.StatusSeeOther)

}
