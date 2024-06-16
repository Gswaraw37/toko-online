package fakers

import (
	"time"

	"github.com/Gswaraw37/toko-online/app/models"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UserFaker(*gorm.DB) *models.User {
	return &models.User{
		ID:            uuid.New().String(),
		FirstName:     faker.FirstName(),
		LastName:      faker.LastName(),
		Email:         faker.Email(),
		Password:      "$2y$10$ZJiR1PAKtfn1MfTpvhocIOkvSui4ynxCFtLqEECAxmByjSFNJxVBK", // password
		RememberToken: "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     gorm.DeletedAt{},
	}
}
