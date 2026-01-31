package services

import "rest-api-go-gin/internal/repositories"

type UserService struct {
	repo *repositories.UserRepository
}
