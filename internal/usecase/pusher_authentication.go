package usecase

import "nextclan/validator-register/mobile-validator-register-service/pkg/pusher"

type PusherAuthenticationUsecase struct {
	pusher pusher.IPusherBeamClient
}

func NewPusherAuthentication(pusherClient pusher.IPusherBeamClient) *PusherAuthenticationUsecase {
	return &PusherAuthenticationUsecase{
		pusherClient,
	}
}

func (p *PusherAuthenticationUsecase) Execute(userId string) (map[string]interface{}, error) {
	token, err := p.pusher.Auth(userId)
	if err != nil {
		return nil, err
	}
	return token, nil
}
