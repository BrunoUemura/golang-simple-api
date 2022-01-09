package controllers

import (
	"github.com/BrunoUemura/golang-simple-api/entities"

	"github.com/gin-gonic/gin"
)

func ShowUsers(c *gin.Context) {
	users := entities.User{}
	users.ID = 1
	users.Name = "Bruno Uemura"
	users.Occupation = "Software Engineer"
	users.Message = "Hello World"

	c.JSON(200, users)
}
