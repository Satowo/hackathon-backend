package usecase

import (
	"hackathon-backend/dao"
	"hackathon-backend/model"
)

func UserSearchUseCase() ([]model.AppUser, error) {
	users, err := dao.UserSearchDao()
	return users, err
}
