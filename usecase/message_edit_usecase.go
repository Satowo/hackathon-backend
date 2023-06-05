package usecase

import (
	"hackathon-backend/dao"
	"hackathon-backend/model"
)

func MessageEditUseCase(messageId string, channelId string, messageContent string) ([]model.Message, error) {
	messages, err := dao.MessageEditDao(messageId, channelId, messageContent)
	return messages, err
}
