package service

import (
	"github.com/cris329/aprendiz-postgres/internal/model"
	"github.com/cris329/aprendiz-postgres/internal/repository"
)

type AprendizService struct {
	repo *repository.AprendizRepository
}

func NewAprendizService(repo *repository.AprendizRepository) *AprendizService {
	return &AprendizService{repo: repo}
}

func (s *AprendizService) CreateAprendiz(a *model.Aprendiz) error {
	return s.repo.Create(a)
}

func (s *AprendizService) GetAll() ([]model.Aprendiz, error) {
	return s.repo.GetAll()
}

func (s *AprendizService) GetByID(id uint) (*model.Aprendiz, error) {
	return s.repo.GetByID(id)
}

func (s *AprendizService) Update(id uint, a *model.Aprendiz) error {
	return s.repo.Update(id, a)
}

func (s *AprendizService) Delete(id uint) error {
	return s.repo.Delete(id)
}
