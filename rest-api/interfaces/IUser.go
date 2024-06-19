package interfaces

import "rest-api/models"

type IUser interface {
	Login(auth *models.User) (*models.User, error)
	Register(auth *models.User) error
	FetchProfile(id string) (*models.User, error)
}
