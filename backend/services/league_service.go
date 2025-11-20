package services

import (
	"backend/entities"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

var (
	ErrLeagueNotFound = errors.New("league not found")
	ErrUnauthorized   = errors.New("you don't have permission to access this league")
)

type LeagueService struct {
	db *gorm.DB
}

func NewLeagueService(db *gorm.DB) *LeagueService {
	return &LeagueService{db: db}
}

func (s *LeagueService) GetLeaguesByOwner(userID uint) ([]entities.League, error) {
	var leagues []entities.League
	if err := s.db.Where("owner_id = ?", userID).Find(&leagues).Error; err != nil {
		return nil, fmt.Errorf("fetching leagues: %w", err)
	}
	return leagues, nil
}

func (s *LeagueService) GetLeagueByID(leagueID, userID uint) (*entities.League, error) {
	var league entities.League
	if err := s.db.First(&league, leagueID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLeagueNotFound
		}
		return nil, fmt.Errorf("fetching league: %w", err)
	}

	if league.OwnerID != userID {
		return nil, ErrUnauthorized
	}

	return &league, nil
}

func (s *LeagueService) CreateLeague(userID uint, name, description, discipline string, playersPerTeam int, hasDraws bool, start, end time.Time) (*entities.League, error) {
	league := &entities.League{
		Name:           name,
		Description:    description,
		PlayersPerTeam: playersPerTeam,
		Discipline:     discipline,
		OwnerID:        userID,
		HasDraws:       hasDraws,
		Start:          start,
		End:            end,
	}

	if err := s.db.Create(league).Error; err != nil {
		return nil, fmt.Errorf("creating league: %w", err)
	}

	return league, nil
}

func (s *LeagueService) UpdateLeague(leagueID, userID uint, name, description string) (*entities.League, error) {
	var league entities.League
	if err := s.db.First(&league, leagueID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrLeagueNotFound
		}
		return nil, fmt.Errorf("fetching league: %w", err)
	}

	if league.OwnerID != userID {
		return nil, ErrUnauthorized
	}

	league.Name = name
	league.Description = description

	if err := s.db.Save(&league).Error; err != nil {
		return nil, fmt.Errorf("updating league: %w", err)
	}

	return &league, nil
}

func (s *LeagueService) DeleteLeague(leagueID, userID uint) error {
	var league entities.League
	if err := s.db.First(&league, leagueID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrLeagueNotFound
		}
		return fmt.Errorf("fetching league: %w", err)
	}

	if league.OwnerID != userID {
		return ErrUnauthorized
	}

	if err := s.db.Delete(&entities.League{}, leagueID).Error; err != nil {
		return fmt.Errorf("deleting league: %w", err)
	}

	return nil
}
