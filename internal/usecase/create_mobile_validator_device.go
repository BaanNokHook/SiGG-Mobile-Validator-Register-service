package usecase

import (
	"context"
	"nextclan/validator-register/mobile-validator-register-service/internal/entity"
	"nextclan/validator-register/mobile-validator-register-service/internal/usecase/repo"
	"nextclan/validator-register/mobile-validator-register-service/pkg/logger"
	"nextclan/validator-register/mobile-validator-register-service/pkg/redis"
	"time"
)

//Use cases include:
/*

 */

//Receive Raw Txn
//Publish to RabbitMQ

type CreateMobileValidatorDeviceUsecase struct {
	log              logger.Interface
	deviceRepository repo.DeviceRepository
	redisClient      redis.RedisInterface
}

func NewCreateMobileValidatorDevice(l logger.Interface, r repo.DeviceRepository, rs redis.RedisInterface) *CreateMobileValidatorDeviceUsecase {
	return &CreateMobileValidatorDeviceUsecase{log: l, deviceRepository: r, redisClient: rs}
}

//Receive new raw txn from clients
func (u *CreateMobileValidatorDeviceUsecase) Execute(device *entity.Device) error {

	currentTime := time.Now().Unix()   
	device.LastestSync = currentTime
	device.Created = currentTime

	_, err := u.deviceRepository.InsertOne(context.Background(), device)
	if err != nil {
		u.log.Error(err)
		panic(err)
	}
	err = u.redisClient.Set(device.PublicKey, device.UserId)
	if err != nil {
		u.log.Error(err)
		panic(err)
	}
	return nil
}
