package main

import (
	"net/http"
	"prj/routes"
	"prj/utils"
)
func main() {
	r := routes.NewRouter()
	utils.LoadTemplates("views/*.html")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}