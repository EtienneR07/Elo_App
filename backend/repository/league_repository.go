package repository

import (
	"backend/models"

	"gorm.io/gorm"
)

type LeagueRepository interface {
	Create(league *models.League) error
	FindByID(id uint) (*models.League, error)
	FindByUserID(userID uint) ([]models.League, error)
	Update(league *models.League) error
	Delete(id uint) error
}

type leagueRepository struct {
	db *gorm.DB
}

func NewLeagueRepository(db *gorm.DB) LeagueRepository {
	return &leagueRepository{db: db}
}

func (r *leagueRepository) Create(league *models.League) error {
	return r.db.Create(league).Error
}

func (r *leagueRepository) FindByID(id uint) (*models.League, error) {
	var league models.League
	err := r.db.First(&league, id).Error
	if err != nil {
		return nil, err
	}
	return &league, nil
}

func (r *leagueRepository) FindByUserID(userID uint) ([]models.League, error) {
	var leagues []models.League
	err := r.db.Where("owner_id = ?", userID).Find(&leagues).Error
	return leagues, err
}

func (r *leagueRepository) Update(league *models.League) error {
	return r.db.Save(league).Error
}

func (r *leagueRepository) Delete(id uint) error {
	return r.db.Delete(&models.League{}, id).Error
}
