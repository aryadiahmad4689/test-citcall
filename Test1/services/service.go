package services

import (
	"github.com/aryadiahmad4689/test-citcall/entity"
	"github.com/aryadiahmad4689/test-citcall/repository"
)

type countryService struct {
	postRepository repository.Country
}

func NewCountryService(postRepository repository.Country) CountryService {
	return &countryService{
		postRepository: postRepository,
	}
}

func (c *countryService) GetAll() ([]entity.Country, error) {
	return c.postRepository.GetAll()
}
