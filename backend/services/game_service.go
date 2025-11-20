package services

import (
	"backend/entities"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var (
	ErrGameNotFound    = errors.New("game not found")
	ErrNotLeagueMember = errors.New("you are not a member of this league")
)

type GameService struct {
	db *gorm.DB
}

func NewGameService(db *gorm.DB) *GameService {
	return &GameService{db: db}
}

func (s *GameService) GetGameByID(gameID, leagueID, userID uint) (*entities.Game, error) {
	var game entities.Game
	if err := s.db.First(&game, gameID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrGameNotFound
		}
		return nil, fmt.Errorf("fetching game: %w", err)
	}

	if game.LeagueID != leagueID {
		return nil, ErrGameNotFound
	}

	if err := s.verifyLeagueMembership(userID, game.LeagueID); err != nil {
		return nil, err
	}

	return &game, nil
}

func (s *GameService) GetGamesByLeagueId(leagueID, userID uint) ([]entities.Game, error) {
	if err := s.verifyLeagueMembership(userID, leagueID); err != nil {
		return nil, err
	}

	var games []entities.Game
	if err := s.db.Where("league_id = ?", leagueID).Find(&games).Error; err != nil {
		return nil, fmt.Errorf("fetching games: %w", err)
	}

	return games, nil
}

func (s *GameService) CreateGame(leagueId uint, team1Score uint, team2Score uint, team1 []uint, team2 []uint) (*entities.Game, error) {
	game := &entities.Game{
		LeagueID: leagueId,
	}

	if err := s.db.Create(game).Error; err != nil {
		return nil, fmt.Errorf("creating game: %w", err)
	}

	var team1Result, team2Result string
	if team1Score > team2Score {
		team1Result = entities.ResultWin
		team2Result = entities.ResultLoss
	} else if team1Score < team2Score {
		team1Result = entities.ResultLoss
		team2Result = entities.ResultWin
	} else {
		team1Result = entities.ResultDraw
		team2Result = entities.ResultDraw
	}

	team1Players, err := s.fetchPlayerLeagues(leagueId, team1)
	if err != nil {
		return nil, fmt.Errorf("fetching team 1 players: %w", err)
	}

	team2Players, err := s.fetchPlayerLeagues(leagueId, team2)
	if err != nil {
		return nil, fmt.Errorf("fetching team 2 players: %w", err)
	}

	team1AvgElo := calculateAverageElo(team1Players)
	team2AvgElo := calculateAverageElo(team2Players)

	team1EloChange := int(CalculateEloChange(team1AvgElo, team2AvgElo, team1Result))
	team2EloChange := int(CalculateEloChange(team2AvgElo, team1AvgElo, team2Result))

	if err := s.createGameParticipants(game.ID, team1Players, 1, team1Result, team1EloChange); err != nil {
		return nil, err
	}

	if err := s.createGameParticipants(game.ID, team2Players, 2, team2Result, team2EloChange); err != nil {
		return nil, err
	}

	return game, nil
}

func (s *GameService) createGameParticipants(
	gameID uint,
	playerLeagues []*entities.PlayerLeague,
	teamNumber int,
	result string,
	eloChange int) error {

	for _, playerLeague := range playerLeagues {
		newElo := uint(int(playerLeague.Elo) + eloChange)

		participant := &entities.GameParticipant{
			GameID:    gameID,
			PlayerID:  playerLeague.ID,
			Team:      teamNumber,
			EloBefore: playerLeague.Elo,
			EloAfter:  newElo,
			EloChange: eloChange,
			Result:    result,
		}
		if err := s.db.Create(participant).Error; err != nil {
			return fmt.Errorf("creating game participant: %w", err)
		}

		playerLeague.Elo = newElo
		switch result {
		case entities.ResultWin:
			playerLeague.Wins++
		case entities.ResultLoss:
			playerLeague.Losses++
		case entities.ResultDraw:
			playerLeague.Draws++
		}

		if err := s.db.Save(playerLeague).Error; err != nil {
			return fmt.Errorf("updating player stats: %w", err)
		}
	}
	return nil
}

func (s *GameService) fetchPlayerLeagues(leagueID uint, playerIDs []uint) ([]*entities.PlayerLeague, error) {
	if len(playerIDs) == 0 {
		return nil, errors.New("team cannot be empty")
	}

	var playerLeagues []*entities.PlayerLeague
	if err := s.db.Where("user_id IN ? AND league_id = ?", playerIDs, leagueID).Find(&playerLeagues).Error; err != nil {
		return nil, fmt.Errorf("fetching player leagues: %w", err)
	}

	if len(playerLeagues) != len(playerIDs) {
		return nil, fmt.Errorf("expected %d players but found %d", len(playerIDs), len(playerLeagues))
	}

	return playerLeagues, nil
}

func calculateAverageElo(playerLeagues []*entities.PlayerLeague) float64 {
	if len(playerLeagues) == 0 {
		return 0
	}

	var totalElo uint
	for _, pl := range playerLeagues {
		totalElo += pl.Elo
	}

	return float64(totalElo) / float64(len(playerLeagues))
}

func (s *GameService) verifyLeagueMembership(userID, leagueID uint) error {
	var count int64
	s.db.Model(&entities.PlayerLeague{}).
		Where("user_id = ? AND league_id = ?", userID, leagueID).
		Count(&count)

	if count == 0 {
		return ErrNotLeagueMember
	}
	return nil
}
