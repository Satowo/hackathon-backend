package usecase

import (
	"hackathon-backend/dao"
	"hackathon-backend/model"
)

func MessageSearchUseCase(channelId string) ([]model.Message, error) {
	messages, err := dao.MessageSearchDao(channelId)
	return messages, err
}
