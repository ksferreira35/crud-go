package service

import (
	"github.com/ksferreira35/crud-go/src/config/logger"
	rest_err "github.com/ksferreira35/crud-go/src/config/res_err"
	"github.com/ksferreira35/crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUser(
	userId string, 
	userDomain model.UserDomainInterface,
)	*rest_err.RestErr {
	logger.Info("Init updateUser model.",
		zap.String("journey", "updateUser"))


	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "updateUser"))
		return nil
	}

	logger.Info(
		"updateUser service executed successfully",
		zap.String("userId", userId), 
		zap.String("journey", "updateUser"))
	return nil
}