package models

import (
	"errors"


	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint64
	FirstName string
	LastName  string
	Email     string
	Password  string
	Status    string
}

func NewUser(user User) (bool, error) {
	con := Connect()
	defer con.Close()
	sql := "insert into users (firstname, lastname, email, password) values ($1, $2, $3, $4)"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	hash, err := Hash(user.Password)
	if err != nil {
		return false, err
	}
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, hash)
	if err != nil {
		return false, err
	}
	return true, err
}

func GetUserByEmail(email string) (User, error) {
	con := Connect()
	defer con.Close()
	sql := "select * from users where email = $1"
	rs, err := con.Query(sql, email)
	if err != nil {
		return User{}, err
	}
	defer rs.Close()
	var user User
	if rs.Next() {
		err := rs.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.Status)
		if err != nil {
			return User{}, err
		}
	}
	return user, nil
}

func CheckEmail(email string) ( error) {
	con := Connect()
	defer con.Close()
	sql := " select count(email) from users where email = $1"
	rs, err := con.Query(sql, email)
	if err != nil {
		return err
	}
	defer rs.Close()
	var count int64
	if rs.Next() {
		err := rs.Scan(&count)
		if err != nil {
			return  err
		}
	}
	if count > 0 {
		return  errors.New("Email is used")
	}
	return nil
}



func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}