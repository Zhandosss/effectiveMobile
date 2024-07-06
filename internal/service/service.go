package service

import (
	"context"
	"effectiveMobileTestProblem/internal/entity"
	"effectiveMobileTestProblem/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (string, error)
	GetUserById(ctx context.Context, id string) (*entity.UserDB, error)
	GetUserByPassport(ctx context.Context, passport string) (*entity.UserDB, error)
	GetUsers(ctx context.Context) ([]*entity.UserDB, error)
	DeleteUserById(ctx context.Context, id string) error
	DeleteUserByPassport(ctx context.Context, passport string) error
	UpdateUserById(ctx context.Context, id string, user *model.User) error
	UpdateUserByPassport(ctx context.Context, passport string, user *model.User) error
}

type WorkRepository interface {
	NewWork(ctx context.Context, work *model.Work) (string, error)
	StopWork(ctx context.Context, id string) error
	GetWorkById(ctx context.Context, id string) (*entity.WorkDB, error)
	GetWorks(ctx context.Context, user string) ([]*entity.WorkDB, error)
}

type UserService struct {
	userRepository UserRepository
}

func NewUser(userRepo UserRepository) *UserService {
	return &UserService{
		userRepository: userRepo,
	}
}

type WorkService struct {
	WorkRepository WorkRepository
}

func NewWork(workRepo WorkRepository) *WorkService {
	return &WorkService{
		WorkRepository: workRepo,
	}
}
