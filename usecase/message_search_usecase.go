package usecase

import (
	"hackathon-backend/dao"
	"hackathon-backend/model"
)

func MessageSearchUseCase() ([]model.Message, error) {
	messages, err := dao.MessageSearchDao()
	return messages, err
}
