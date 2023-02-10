package routes

import (
	"net/http"
	"prj/models"
	"prj/sessions"
	"prj/utils"
)

func adminHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")
	idUser, _:= session.Values["USERID"].(uint64)

	todoList, _:= models.GetTodoList(idUser)
	utils.ExcuteTemplate(w, "admin.html", struct{
		TodoList []models.Todo
	}{
		TodoList : todoList,
	})
}