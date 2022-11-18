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

type FindByIdMobileValidatorDeviceStatus struct {
	log              logger.Interface
	deviceRepository repo.DeviceRepository
}

func NewFindMobileValidatorDeviceStatus(l logger.Interface, r repo.DeviceRepository) *FindByIdMobileValidatorDeviceStatus {
	return &FindByIdMobileValidatorDeviceStatus{log: l, deviceRepository: r}
}

// Receive new verified txn from clients
func (u *FindByIdMobileValidatorDeviceStatus) Execute(device *entity.Device) error {
	currentTime := time.Now().Unix()
	u.deviceRepository.FindDeviceByLastestSyncGte(context.TODO(), currentTime, device.ID.Timestamp().Unix())
	return nil
}
