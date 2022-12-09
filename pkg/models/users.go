package models

import (
	"github.com/google/uuid"
	"github.com/usblco/polarisb-authn-go/internal/pkg/internalModels"
)

type PolarisbUser struct {
	Id           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	Passwordhash string
	Fullname     string `json:"fullname"`
	CreateDate   string `json:"createDate"`
	Role         int    `json:"role"`
}

func (m PolarisbUser) CheckPassword(password string) bool {

	passwordObj := &internalModels.PasswordObject{}
	err := passwordObj.DecodeFromBase64(m.Passwordhash)
	if err != nil {
		return false
	}

	// Check password
	err = passwordObj.ComparePasswordHash(password)
	if err != nil {
		return false
	}

	return true
}
