package usecase

import (
	"hackathon-backend/dao"
	"hackathon-backend/model"
)

func MessageDeleteUseCase(messageId string, channelId string) ([]model.Message, error) {
	messages, err := dao.MessageDeleteDao(messageId, channelId)
	return messages, err
}
