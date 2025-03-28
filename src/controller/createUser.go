package controller

import (
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
		zap.String("journey", "createUser "),
	)
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

	domainResult, err := uc.service.CreateUserServices(domain)
	if err != nil {
		logger.Error(
			"Error to trying to call CreateUser service",
			err,
			zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"CreateUser controller executed successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "createUser "))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))
}
