package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Device struct {
	ID                         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserId                     string             `bson:"userId" json:"userId"`
	DeviceId                   string             `bson:"deviceId" json:"deviceId"`
	PublicKey                  string             `bson:"publicKey" json:"publicKey"`
	LastestSync                int64              `bson:"lastestSync" json:"lastestSync"`
	Created                    int64              `bson:"created" json:"created"`
	LastestValidateTransaction int64              `bson:"lastestValidateTransaction" json:"lastestValidateTransaction"`
}
