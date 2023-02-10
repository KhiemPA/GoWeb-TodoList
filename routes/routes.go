package routes

import (
	"net/http"
	"prj/middleware"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/", redirectHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/register", registerHandler)
	r.HandleFunc("/logout", logOutHandler)
	r.HandleFunc("/admin", middleware.AuthRequired(adminHandler)).Methods("GET")
	r.HandleFunc("/create-task", middleware.AuthRequired(creatTaskHandler))
	r.HandleFunc("/task-edit", middleware.AuthRequired(updateTaskHandler))
	r.HandleFunc("/task-delete", middleware.AuthRequired(deleteTaskHandler))
	r.HandleFunc("/checked", middleware.AuthRequired(setDoneStatusForTaskHandler))
	r.HandleFunc("/unchecked", middleware.AuthRequired(setUndoStatusForTaskHandler))
	fileServer := http.FileServer(http.Dir("views/assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))
	return r
}