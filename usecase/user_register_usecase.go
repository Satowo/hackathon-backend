package usecase

import (
	"crypto/rand"
	"github.com/oklog/ulid/v2"
	"hackathon-backend/dao"
	"hackathon-backend/model"
)

func UserRegisterUseCase(name string, age int) ([]model.User, error) {
	// IDを生成
	entropy := ulid.Monotonic(rand.Reader, 0)
	id := ulid.MustNew(ulid.Now(), entropy).String()

	users, err := dao.UserRegisterDao(id, name, age)
	return users, err
}
