package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ksferreira35/crud-go/src/config/validation"
	"github.com/ksferreira35/crud-go/src/controller/model/request"
)

func CreateUser (c *gin.Context) {
	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshall object, error=%s\n", err.Error())
		errRest := validation.ValidateUserError(err)
		
		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)
}
