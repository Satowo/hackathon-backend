package usecase

import (
	"hackathon-backend/dao"
	"hackathon-backend/model"
)

func ChannelMemberSearchUseCase(channelId string) ([]model.AppUser, error) {
	members, err := dao.ChannelMemberSearchDao(channelId)
	if err != nil {
		return nil, err
	}
	return members, nil
}
