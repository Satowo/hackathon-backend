package usecase

import (
	"crypto/rand"
	"github.com/oklog/ulid/v2"
	"hackathon-backend/dao"
	"hackathon-backend/model"
)

func ChannelLeaveUseCase(channelId string, userId string) (model.UserInfo, error) {
	// IDを生成
	entropy := ulid.Monotonic(rand.Reader, 0)
	Id := ulid.MustNew(ulid.Now(), entropy).String()

	userInfo, err := dao.ChannelLeaveDao(Id, channelId, userId)
	return userInfo, err
}
