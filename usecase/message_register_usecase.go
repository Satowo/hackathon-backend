package usecase

import (
	"crypto/rand"
	"github.com/oklog/ulid/v2"
	"hackathon-backend/dao"
	"hackathon-backend/model"
)

func MessageRegisterUseCase(userId string, channelId string, messageContent string) ([]model.Message, error) {
	// IDを生成
	entropy := ulid.Monotonic(rand.Reader, 0)
	messageId := ulid.MustNew(ulid.Now(), entropy).String()

	messages, err := dao.MessageRegisterDao(messageId, userId, channelId, messageContent)
	return messages, err
}
