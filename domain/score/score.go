package score

import (
	"github.com/qinsheng99/go-py/api/score_api"
)

type Score interface {
	Score1(score_api.Score) ([]byte, error)
	Score2(score_api.Score) ([]byte, error)
}
