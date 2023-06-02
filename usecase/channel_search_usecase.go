package usecase

import (
	"hackathon-backend/dao"
	"hackathon-backend/model"
)

func ChannelSearchUseCase() ([]model.Channel, error) {
	channels, err := dao.ChannelSearchDao()
	return channels, err
}
