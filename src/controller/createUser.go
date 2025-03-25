package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ksferreira35/crud-go/src/config/logger"
	"github.com/ksferreira35/crud-go/src/config/validation"
	"github.com/ksferreira35/crud-go/src/controller/model/request"
	"github.com/ksferreira35/crud-go/src/model"
	"github.com/ksferreira35/crud-go/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser "))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser "))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	domainResult, err := uc.service.CreateUser(domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	userString := fmt.Sprintf("%s, %s, %d", domain.GetName(), domain.GetEmail(), domain.GetAge())

	logger.Info("User created successfully",
		zap.String("journey", "createUser "),
		zap.String("user", userString))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))
}
