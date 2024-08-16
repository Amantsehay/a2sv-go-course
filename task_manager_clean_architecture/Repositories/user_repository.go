package Repositories

import (
    "task_manager_clean_architecture/Domain"
)

type UserRepository interface {
    CreateUser(username, password, role string) (Domain.User, error)
    AuthenticateUser(username, password string) (Domain.User, error)
    PromoteUser(userID string) error
}
