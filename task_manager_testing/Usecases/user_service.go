package Usecases

import (
    "task_manager_clean_architecture/Domain"
    "task_manager_clean_architecture/Repositories"
)

type UserUsecases struct {
    repo Repositories.UserRepository
}

func NewUserUsecases(repo Repositories.UserRepository) *UserUsecases {
    return &UserUsecases{repo: repo}
}

func (uc *UserUsecases) CreateUser(username, password, role string) (*Domain.User, error) {
    user, err := uc.repo.CreateUser(username, password, role)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (uc *UserUsecases) AuthenticateUser(username, password string) (*Domain.User, error) {
    user, err := uc.repo.AuthenticateUser(username, password)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (uc *UserUsecases) PromoteUser(userID string) error {
    return uc.repo.PromoteUser(userID)
}
