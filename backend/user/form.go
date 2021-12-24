package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenDetails struct {
	AccessToken        string `json:"access_token"`
	AccessTokenExpires int64  `json:"access_token_expires"`
	AccessTokenUUID    string `json:"access_token_uuid"`
}

type UserSignup struct {
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	CreateDate time.Time `json:"create_date"`
}

type UserInfo struct {
	UserId     primitive.ObjectID `bson:"_id" json:"id,omitempty"` 
	Email      string    `json:"email"`
	Password   string    `json:"password"`
}

type UserLogin struct {
	Email      string    `json:"email"`
	Password   string    `json:"password"`
}

type UserRequest struct {
	UserId      primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}