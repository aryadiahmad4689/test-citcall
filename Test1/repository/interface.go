package repository

import (
	"github.com/aryadiahmad4689/test-citcall/entity"
)

type Country interface {
	GetAll() ([]entity.Country, error)
}
