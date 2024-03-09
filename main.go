package main

import (
	"challange1_9Mar/config"
	"challange1_9Mar/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	inDB := &controller.InDB{DB: db}

	router := gin.Default()

	router.GET("person/:id", inDB.GetPerson)
	router.GET("persons", inDB.GetPersons)
	router.POST("person", inDB.CreatePerson)
	router.PUT("person", inDB.UpdatePerson)
	router.DELETE("person/:id", inDB.DeletePerson)
	router.Run(":8080")
}
