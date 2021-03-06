// Code generated by git.repo.services.lenvendo.ru/grade-factor/go-kit-service-generator  REMOVE THIS STRING ON EDIT OR DO NOT EDIT.
//go:generate mockgen -destination service_mock.go -package user  "github.com/kagestonedragon/server/pkg/user Service
package user

import (
	"context"

	_ "github.com/golang/mock/mockgen/model"
)

type Service interface {
	AddUser(context.Context, *AddUserRequest) (*User, error)
	GetUserList(context.Context, *GetUserListRequest) (*GetUserListResponse, error)
	UpdateUser(context.Context, *User) (*UpdateUserResponse, error)
	DeleteUserById(context.Context, *DeleteUserByIdRequest) (*DeleteUserByIdResponse, error)
	GetUserById(context.Context, *GetUserByIdRequest) (*User, error)
}
