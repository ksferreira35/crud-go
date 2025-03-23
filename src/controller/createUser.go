package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ksferreira35/crud-go/src/config/res_err"
)

func CreateUser (c *gin.Context) {

	err := rest_err.NewBadRequestError("Você chamou a rota de forma errada besta")
	c.JSON(err.Code, err)
}