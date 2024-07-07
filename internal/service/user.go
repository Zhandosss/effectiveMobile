package service

import (
	"context"
	"effectiveMobileTestProblem/internal/model"
	"errors"
)

func (s *UserService) CreateUser(ctx context.Context, user *model.User) (string, error) {
	_, err := s.userRepository.GetUserByPassport(ctx, user.PassportSeriesAndNumber)
	if errors.Is(err, model.ErrNotFound) {
		return s.userRepository.CreateUser(ctx, user)
	}
	if err != nil {
		return "", err
	}
	return "", model.ErrAlreadyExists
}

func (s *UserService) GetUserById(ctx context.Context, id string) (*model.User, error) {
	user, err := s.userRepository.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user.ToUser(), nil
}

func (s *UserService) GetUserByPassport(ctx context.Context, passport string) (*model.User, error) {
	user, err := s.userRepository.GetUserByPassport(ctx, passport)
	if err != nil {
		return nil, err
	}
	return user.ToUser(), nil
}

func (s *UserService) GetUsers(ctx context.Context, filterAndPagination *model.FilterAndPagination) ([]*model.User, error) {
	users, err := s.userRepository.GetUsers(ctx, filterAndPagination)
	if err != nil {
		return nil, err
	}
	var result []*model.User
	for _, user := range users {
		result = append(result, user.ToUser())
	}
	return result, nil
}

func (s *UserService) DeleteUserById(ctx context.Context, id string) error {
	return s.userRepository.DeleteUserById(ctx, id)
}

func (s *UserService) DeleteUserByPassport(ctx context.Context, passport string) error {
	return s.userRepository.DeleteUserByPassport(ctx, passport)
}

func (s *UserService) UpdateUserById(ctx context.Context, id string, user *model.User) error {
	return s.userRepository.UpdateUserById(ctx, id, user)
}

func (s *UserService) UpdateUserByPassport(ctx context.Context, passport string, user *model.User) error {
	return s.userRepository.UpdateUserByPassport(ctx, passport, user)
}
