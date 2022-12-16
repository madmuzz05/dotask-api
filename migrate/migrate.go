package main

import (
	"github.com/madmuzz05/dotask-api.git/initialize"
	model "github.com/madmuzz05/dotask-api.git/models"
)

func main() {
	db := initialize.ConnectToDB()
	db.AutoMigrate(&model.Task{})
}
