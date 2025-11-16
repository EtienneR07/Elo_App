package services

import (
	"backend/models"
	"backend/repository"
	"errors"

	"gorm.io/gorm"
)

var (
	ErrLeagueNotFound = errors.New("league not found")
	ErrUnauthorized   = errors.New("unauthorized")
)

// LeagueService handles league business logic
type LeagueService interface {
	CreateLeague(league *models.League) error
	GetLeagueByID(id, userID uint) (*models.League, error)
	GetUserLeagues(userID uint) ([]models.League, error)
	UpdateLeague(id, userID uint, league *models.League) error
	DeleteLeague(id, userID uint) error
}

type leagueService struct {
	leagueRepo repository.LeagueRepository
}

// NewLeagueService creates a new league service instance
func NewLeagueService(leagueRepo repository.LeagueRepository) LeagueService {
	return &leagueService{leagueRepo: leagueRepo}
}

func (s *leagueService) CreateLeague(league *models.League) error {
	return s.leagueRepo.Create(league)
}

func (s *leagueService) GetLeagueByID(id, userID uint) (*models.League, error) {
	league, err := s.leagueRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLeagueNotFound
		}
		return nil, err
	}

	// Check ownership
	if league.OwnerID != userID {
		return nil, ErrUnauthorized
	}

	return league, nil
}

func (s *leagueService) GetUserLeagues(userID uint) ([]models.League, error) {
	return s.leagueRepo.FindByUserID(userID)
}

func (s *leagueService) UpdateLeague(id, userID uint, updatedLeague *models.League) error {
	league, err := s.leagueRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrLeagueNotFound
		}
		return err
	}

	// Check ownership
	if league.OwnerID != userID {
		return ErrUnauthorized
	}

	// Update fields
	league.Name = updatedLeague.Name
	league.Description = updatedLeague.Description
	league.PlayersPerTeam = updatedLeague.PlayersPerTeam
	league.Discipline = updatedLeague.Discipline

	return s.leagueRepo.Update(league)
}

func (s *leagueService) DeleteLeague(id, userID uint) error {
	league, err := s.leagueRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrLeagueNotFound
		}
		return err
	}

	// Check ownership
	if league.OwnerID != userID {
		return ErrUnauthorized
	}

	return s.leagueRepo.Delete(id)
}
