package usecase

import (
	"github.com/AkiOuma/biz-groups/internal/domain/repository"
	"github.com/AkiOuma/biz-groups/internal/domain/service"
)

type Usecase interface {
	Group
}

type uc struct {
	repo repository.Repository
	svc  service.Service
}

var _ Usecase = (*uc)(nil)

// Constructor
func NewUsecase(
	repo repository.Repository,
	svc service.Service,
) Usecase {
	return &uc{
		svc:  svc,
		repo: repo,
	}
}
