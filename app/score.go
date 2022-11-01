package app

import (
	"bytes"
	"encoding/json"
	"github.com/qinsheng99/go-py/api/score_api"
	"github.com/qinsheng99/go-py/domain/score"
)

type scoreService struct {
	s score.Score
}

type ScoreService interface {
	Score1(score_api.Score, *score_api.ScoreRes) error
	Score2(score_api.Score, *score_api.ScoreRes) error
}

func NewScoreService(s score.Score) ScoreService {
	return &scoreService{
		s: s,
	}
}

func (s *scoreService) Score1(col score_api.Score, res *score_api.ScoreRes) error {
	bys, err := s.s.Score1(col)
	if err != nil {
		return err
	}

	err = json.NewDecoder(bytes.NewBuffer(bys)).Decode(res)
	if err != nil {
		return err
	}
	return nil
}

func (s *scoreService) Score2(col score_api.Score, res *score_api.ScoreRes) error {
	bys, err := s.s.Score2(col)
	if err != nil {
		return err
	}

	err = json.NewDecoder(bytes.NewBuffer(bys)).Decode(res)
	if err != nil {
		return err
	}
	return nil
}
