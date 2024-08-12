package mappers

import (
	"project-microservices/dto"
	userServiceProto "project-microservices/user_service/proto"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserVerificationKeyToGrpc(req *dto.PhoneVerificationReq) *userServiceProto.PhoneVerificationReq {
	return &userServiceProto.PhoneVerificationReq{
		Phone: req.Phone,
	}
}
func UserVerificationKeyFromGrpc(res *userServiceProto.PhoneVerificationRes) *dto.PhoneVerificationRes {
	return &dto.PhoneVerificationRes{
		VerificationKey: res.VerificationKey,
	}
}

func UserPhoneCreateToGrpc(req dto.UserCreateReq) *userServiceProto.UserPhoneCreateReq {
	return &userServiceProto.UserPhoneCreateReq{
		Sms:             req.SMS,
		VerificationKey: req.VerificationKey,
	}
}

func UserPhoneCreateFromGrpc(res *userServiceProto.UserPhoneCreateRes) *dto.UserCreateRes {
	return &dto.UserCreateRes{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
		UserType:     int(res.UserType),
	}
}

func GetUserToGrpc(req dto.GetUserReq) *userServiceProto.GetUser {
	return &userServiceProto.GetUser{
		Token: req.Token,
	}
}

func GetUserFromGrpc(res *userServiceProto.User) *dto.User {
	return &dto.User{
		ID:     uuid.FromStringOrNil(res.Id),
		Phone:  res.Phone,
		UserId: res.UserId,
		Email:  &res.Email,
	}
}

func UserUpdateToGrpc(req *dto.UserUpdateReq) *userServiceProto.PostUser {
	return &userServiceProto.PostUser{
		Phone:     req.Phone,
		Username:  req.Username,
		UpdatedAt: timestamppb.Now(),
		UserId:    req.UserId,
		Email:     req.Email,
	}
}

func UserToGrpc(req *dto.User) *userServiceProto.User {
	return &userServiceProto.User{
		UserId: req.UserId,
		Phone:  req.Phone,
		Email:  *req.Email,
		Id:     req.ID.String(),
	}
}

func UserFromGrpc(res *userServiceProto.User) *dto.User {
	return &dto.User{
		ID:     uuid.FromStringOrNil(res.Id),
		Phone:  res.Phone,
		UserId: res.UserId,
		Email:  &res.Email,
	}
}
