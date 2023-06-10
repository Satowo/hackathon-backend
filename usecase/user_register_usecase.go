package usecase

import (
	"crypto/rand"
	"github.com/oklog/ulid/v2"
	"hackathon-backend/dao"
	"hackathon-backend/model"
)

func UserRegisterUseCase(userName string, email string, password string) (model.UserInfo, error) {
	// IDを生成
	entropy := ulid.Monotonic(rand.Reader, 0)
	userId := ulid.MustNew(ulid.Now(), entropy).String()

	userInfo, err := dao.UserRegisterDao(userId, userName, email, password)
	return userInfo, err
}
