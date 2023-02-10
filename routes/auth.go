package routes

import (
	"errors"
	"fmt"
	"net/http"
	"prj/models"
	"prj/sessions"
	"prj/utils"
	"golang.org/x/crypto/bcrypt"
)


var (
	ErrEmtpyField = errors.New("Empty field")
	ErrEmailNotFound = errors.New("Email doesnt exist")
	ErrFirstNameRequired = errors.New("First name is from 1 to 15 char")
	ErrLastNameRequired = errors.New("Last name is from 1 to 20 char")
	ErrPasswordRequired = errors.New("Password is from 1 to 100 char")
	ErrPasswordIncorrect = errors.New("Password incorrect")

)


func ValidateFileds(email, password string) error {
	if len(email) == 0 || len(password) == 0 {
		return ErrEmtpyField
	}
	return nil
}

func CheckDataFromUser(firstname, lastname, password string) error{
	if len(firstname) == 0 || len(firstname) > 15 {
		return ErrFirstNameRequired
	}
	if len(lastname) == 0 || len(lastname) > 20 {
		return ErrLastNameRequired
	}
	if len(password) == 0 || len(password) > 100 {
		return ErrPasswordRequired
	}
	return nil
}


func CheckAccount(email, password string) (models.User, error) {
	if err := ValidateFileds(email, password); err != nil {
		return models.User{} ,err
	}
	user, err := models.GetUserByEmail(email)
	if err != nil {
		return user, err
	}
	if user.Id == 0 {
		return user, ErrEmailNotFound
	}

	// hashed password rồi tới password
	errPassword := bcrypt.CompareHashAndPassword( []byte(user.Password), []byte(password))
	if errPassword != nil {
		return user, ErrPasswordIncorrect

	}
	return user, nil
}




func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		utils.ExcuteTemplate(w, "login.html", nil)
		return
	} else {
		r.ParseForm()
		email := r.PostForm.Get("email")
		password := r.PostForm.Get("password")

		user, err := CheckAccount(email, password)
		
		if err != nil {
			fmt.Println(err)
			var data = ""
			if err != nil {
				data = err.Error()
			} 
			utils.ExcuteTemplate(w, "login.html", struct{
				Data string
			}{
				Data: data,
			})
			return
		} else {
			session, _ := sessions.Store.Get(r, "session")
			session.Values["USERID"] = user.Id
			session.Save(r, w)
			http.Redirect(w, r, "/admin", 302)
			return	
		}
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		utils.ExcuteTemplate(w, "register.html", nil)
		return
	} else {
		r.ParseForm()
		var user models.User
		user.FirstName = r.PostForm.Get("firstname")
		user.LastName = r.PostForm.Get("lastname")
		user.Email = r.PostForm.Get("email")
		user.Password = r.PostForm.Get("password")

		err := models.CheckEmail(user.Email)
		if err != nil {
			utils.ExcuteTemplate(w, "register.html", struct{
				Err string
			}{
				Err : err.Error(),
			})
			return
		}
		err = CheckDataFromUser(user.FirstName, user.LastName, user.Password)
		if err != nil {
			utils.ExcuteTemplate(w, "register.html", struct{
				Err string
			}{
				Err : err.Error(),
			})
			return
		}

		_, err = models.NewUser(user)
		if err != nil {
			utils.ExcuteTemplate(w, "register.html", struct{
				Err string
			}{
				Err : err.Error(),
			})
			return
		}

		http.Redirect(w, r, "/login", 302)
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", 302)
	return
}

func logOutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		session, _ := sessions.Store.Get(r, "session")
		delete(session.Values, "USERID")
		session.Save(r, w)
		http.Redirect(w, r, "/login", 302)
		return
	}
	return
}