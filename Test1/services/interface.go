package services

import "github.com/aryadiahmad4689/test-citcall/entity"

type CountryService interface {
	GetAll() ([]entity.Country, error)
}
