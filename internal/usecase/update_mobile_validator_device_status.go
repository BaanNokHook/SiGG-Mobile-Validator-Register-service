package usecase

import (
	"context"
	"nextclan/validator-register/mobile-validator-register-service/internal/entity"
	"nextclan/validator-register/mobile-validator-register-service/internal/usecase/repo"
	"nextclan/validator-register/mobile-validator-register-service/pkg/logger"
	"time"
)

//Use cases include:
/*

 */

//Receive Verified Txn
//Publish to RabbitMQ

type UpdateMobileValidatorDeviceStatusUseCase struct {
	log              logger.Interface
	deviceRepository repo.DeviceRepository
}

func NewUpdateMobileValidatorDeviceStatus(l logger.Interface, r repo.DeviceRepository) *UpdateMobileValidatorDeviceStatusUseCase {
	return &UpdateMobileValidatorDeviceStatusUseCase{log: l, deviceRepository: r}
}

//Receive new verified txn from clients
func (u *UpdateMobileValidatorDeviceStatusUseCase) Execute(device *entity.Device) error {
	currentTime := time.Now().Unix()
	u.deviceRepository.UpdateLastestSyncByDeviceId(context.TODO(), currentTime, device)
	return nil
}
