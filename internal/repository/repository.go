package repository

import (
	"github.com/cris329/aprendiz-postgres/internal/model"

	"gorm.io/gorm"
)

type AprendizRepository struct {
	db *gorm.DB
}

func NewAprendizRepository(db *gorm.DB) *AprendizRepository {
	return &AprendizRepository{db: db}
}

func (r *AprendizRepository) Create(a *model.Aprendiz) error {
	return r.db.Create(a).Error
}

func (r *AprendizRepository) GetAll() ([]model.Aprendiz, error) {
	var aprendices []model.Aprendiz
	err := r.db.Find(&aprendices).Error
	return aprendices, err
}

func (r *AprendizRepository) GetByID(id uint) (*model.Aprendiz, error) {
	var a model.Aprendiz
	err := r.db.First(&a, id).Error
	return &a, err
}

func (r *AprendizRepository) Update(id uint, a *model.Aprendiz) error {
	return r.db.Model(&model.Aprendiz{}).Where("id = ?", id).Updates(a).Error
}

func (r *AprendizRepository) Delete(id uint) error {
	return r.db.Delete(&model.Aprendiz{}, id).Error
}
