package usecase

import (
	"hackathon-backend/dao"
	"hackathon-backend/model"
)

func UserInfoUseCase(email string) (model.UserInfo, error) {
	userInfo, err := dao.UserInfoDao(email)
	return userInfo, err
}
