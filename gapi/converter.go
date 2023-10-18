package gapi

import (
	db "github.com/zaenalarifin12/simplebank/db/sqlc"
	"github.com/zaenalarifin12/simplebank/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(users db.Users) *pb.User {
	return &pb.User{
		Username:         users.Username,
		FullName:         users.FullName,
		Email:            users.Email,
		PasswordChangeAt: timestamppb.New(users.PasswordChangeAt),
		CreatedAt:        timestamppb.New(users.CreatedAt),
	}
}
