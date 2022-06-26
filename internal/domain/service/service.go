package service

import "github.com/AkiOuma/biz-groups/internal/domain/repository"

type Service interface {
	Group
}

type svc struct {
	repo repository.Repository
}

var _ Service = (*svc)(nil)

func NewService(repo repository.Repository) Service {
	return &svc{repo: repo}
}
