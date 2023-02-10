package models

import (
	"fmt"
	"time"
)

type Todo struct {
	Id      uint64
	Message string
	Status 	bool
	Created string
	Updated string
	IdUser  uint64
}

func NewTodo(message string, idUser uint64) (bool, error) {
	con := Connect()
	defer con.Close()
	var todo Todo
	todo.Message = message
	now := time.Now()
	todo.Created = now.String()
	todo.IdUser = idUser
	sql := "insert into todo (message, status, created, updated, idUser) values($1, $2, $3, $4, $5)"
	stmt, err := con.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	defer stmt.Close()

	_,err = stmt.Exec(todo.Message, todo.Status, todo.Created, todo.Updated, todo.IdUser)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}


func SetStatusForTask(id uint64, status bool) error{
	con := Connect()
	defer con.Close()
	sql := "update todo set status=$1, updated=$2 where id=$3"
	stmt, err := con.Prepare(sql)
	if err != nil {
		fmt.Println(err)

		return  err
	}
	defer stmt.Close()
	now := time.Now()
	_,err = stmt.Exec(status, now.String(), id)
	if err != nil {
		fmt.Println(err)

		return err
	}
	return nil
}

func GetTaskById(id uint64) (Todo, error) {
	con := Connect()
	defer con.Close()
	sql := "select * from todo where id=$1"

	rs, err := con.Query(sql, id)
	if err != nil {
		fmt.Println(err)

		return Todo{}, err
	}
	defer rs.Close()
	var todo Todo
	for rs.Next() {
		err := rs.Scan(&todo.Id, &todo.Message, &todo.Status, &todo.Created, &todo.Updated, &todo.IdUser)
		if err != nil {
			fmt.Println(err)

			return Todo{}, err
		}
	}
	return todo, nil
}


func GetTodoList(idUser uint64) ([]Todo, error){
	con := Connect()
	sql := "select * from todo where idUser=$1 order by id desc"
	rs, err := con.Query(sql, idUser)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rs.Close()
	todoList := []Todo{}
	for rs.Next() {
		todo := Todo{}
		err := rs.Scan(&todo.Id, &todo.Message, &todo.Status, &todo.Created, &todo.Updated, &todo.IdUser)
		if err != nil {
			return nil, err
		}
		todoList = append(todoList, todo)
	}
	return todoList, nil
}

func DeleteTask(id uint64) (error) {
	con := Connect()
	defer con.Close()
	sql := "delete from todo where id=$1"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return err
	}
	_,err = stmt.Query(id)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTask(id uint64, message string) (bool,error) {
	con := Connect()
	defer con.Close()
	sql := "update todo set message=$1, updated=$2 where id=$3"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	now := time.Now()
	_,err = stmt.Exec(message, now.String(), id)
	if err != nil {
		return false, err
	}
	return true, nil
}