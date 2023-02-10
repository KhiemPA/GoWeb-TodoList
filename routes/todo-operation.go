package routes

import (
	"fmt"
	"net/http"
	"prj/models"
	"prj/sessions"
	"prj/utils"
	"strconv"
)


func creatTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		utils.ExcuteTemplate(w, "create-task.html", nil)
		return
	}
	r.ParseForm()
	message := r.PostForm.Get("message")
	session, _ := sessions.Store.Get(r, "session")
	idUser, _:= session.Values["USERID"].(uint64)
	models.NewTodo(message, idUser)
	http.Redirect(w, r, "/admin", 302)
	return
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query()
	id, _ := strconv.ParseUint(key.Get("taskId"), 10, 64)
	models.DeleteTask(id)
	http.Redirect(w, r, "/admin",302)
	return
}

func setDoneStatusForTaskHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	id, _ := strconv.ParseUint(keys.Get("taskId"), 10, 64)
	models.SetStatusForTask(id, true)
	http.Redirect(w, r, "/admin", 302)
	return
}

func setUndoStatusForTaskHandler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	id, _ := strconv.ParseUint(keys.Get("taskId"), 10, 64)
	models.SetStatusForTask(id, false)
	http.Redirect(w, r, "/admin", 302)
	return
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		keys := r.URL.Query()
		id, _ := strconv.ParseUint(keys.Get("taskId"), 10, 64)
		task, err := models.GetTaskById(id)
		if err != nil {
			http.Redirect(w, r, "/admin", 302)
			return
			
		}
		utils.ExcuteTemplate(w, "edit-task.html", task)
		return
		
	}
	r.ParseForm()
	id, _ := strconv.ParseUint(r.PostForm.Get("id"), 10, 64)
	newMessage := r.PostForm.Get("message")
	ok, err := models.UpdateTask(id, newMessage)
	if !ok {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/admin", 302)
	return
	
}



