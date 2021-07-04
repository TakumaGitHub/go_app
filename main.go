package main

import (
	"fmt"
	"main/app/controllers"
	"main/app/models"
	// "todo_app/config"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

}
