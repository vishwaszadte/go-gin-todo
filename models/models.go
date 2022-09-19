package models

import (
	"github.com/vishwaszadte/go-gin-todo/database"
	"gorm.io/gorm"
)

var db *gorm.DB

type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

func init() {
	database.Connect()
	db = database.GetDB()
	db.AutoMigrate(&Todo{})
}

func GetTodos(todos *[]Todo) (err error) {
	if err := db.Find(todos).Error; err != nil {
		return err
	}

	return nil
}

func GetTodo(todo *Todo, id string) (err error) {
	if err := db.Where("id = ?", id).First(todo).Error; err != nil {
		return err
	}

	return nil
}

func CreateTodo(todo *Todo) (err error) {
	if err := db.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTodo(todo *Todo, id string) (err error) {
	db.Where("id = ?", id).Delete(todo)
	return nil
}

func UpdateTodo(todo *Todo, id string) (err error) {
	db.Save(todo)
	return nil
}
