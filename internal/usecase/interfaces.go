package usecase

import (
	"nextclan/validator-register/mobile-validator-register-service/internal/entity"
)

type (
	CreateMobileValidatorDevice interface {
		Execute(*entity.Device) error
	}

	UpdateMobileValidatorDeviceStatus interface {
		Execute(*entity.Device) error
	}

	PusherAuthentication interface {
		Execute(userId string) (map[string]interface{}, error)
	}
)
