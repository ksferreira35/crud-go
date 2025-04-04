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

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init loginUser controller",
		zap.String("journey", "loginUser "),
	)
	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "loginUser "))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)

	domainResult, token, err := uc.service.LoginUserServices(domain)
	if err != nil {
		logger.Error(
			"Error to trying to call loginUser service",
			err,
			zap.String("journey", "loginUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"loginUser controller executed successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "loginUser "))

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,	
	))
}
