package main

import (
	"crud_go_native/config"
	categorycontroller "crud_go_native/controllers/categoryController"
	homecontroller "crud_go_native/controllers/homeController"
	productcontroller "crud_go_native/controllers/productController"
	registercontroller "crud_go_native/controllers/registerController"
	"crud_go_native/middleware"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	config.ConnectDB()

	mux := http.NewServeMux()

	// home page
	mux.HandleFunc("/", homecontroller.Welcome)

	// categories
	mux.HandleFunc("/categories", categorycontroller.Index)
	mux.HandleFunc("/categories/add", categorycontroller.Add)
	mux.HandleFunc("/categories/update", categorycontroller.Update)
	mux.HandleFunc("/categories/delete", categorycontroller.Delete)

	// products
	mux.HandleFunc("/products", productcontroller.Index)
	mux.HandleFunc("/products/add", productcontroller.Add)
	mux.HandleFunc("/products/detail", productcontroller.Detail)
	mux.HandleFunc("/products/update", productcontroller.Update)
	mux.HandleFunc("/products/delete", productcontroller.Delete)

	// register
	mux.HandleFunc("/register/signup", registercontroller.SignUp)
	mux.HandleFunc("/register/login", registercontroller.LogIn)

	logger.Info("server running at port :8000")
	middle := &middleware.LogMiddleware{
		Handler: mux,
	}
	http.ListenAndServe(":8000", middle)
}
