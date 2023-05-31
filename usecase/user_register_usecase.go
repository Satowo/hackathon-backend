package usecase

import (
	"crypto/rand"
	"github.com/oklog/ulid/v2"
	"hackathon-backend/dao"
)

func UserRegisterUseCase(name string, email string, pwd string) error {
	// IDを生成
	entropy := ulid.Monotonic(rand.Reader, 0)
	id := ulid.MustNew(ulid.Now(), entropy).String()

	err := dao.UserRegisterDao(id, name, email, pwd)
	return err
}
