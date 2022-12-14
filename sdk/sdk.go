package sdk

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/opensourceways/community-robot-lib/utils"
	"github.com/qinsheng99/go-py/api/score_api"
	"net/http"
	"strings"
)

type Calculate = score_api.Score

func NewCalculateEvaluate(endpoint string) CalculateEvaluate {
	return CalculateEvaluate{
		endpoint: strings.TrimSuffix(endpoint, "/"),
		cli:      utils.NewHttpClient(3),
	}
}

func (t CalculateEvaluate) calculateURL() string {
	return fmt.Sprintf("%s/v1/calculate", t.endpoint)
}

func (t CalculateEvaluate) evaluateURL() string {
	return fmt.Sprintf("%s/v1/evaluate", t.endpoint)
}

func (t CalculateEvaluate) Calculate(opt *Calculate) error {
	payload, err := utils.JsonMarshal(opt)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, t.calculateURL(), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	return t.forwardTo(req, nil)
}

func (t CalculateEvaluate) Evaluate(opt *Calculate) (_ error, score float64) {
	payload, err := utils.JsonMarshal(opt)
	if err != nil {
		return err, 0
	}

	req, err := http.NewRequest(http.MethodPost, t.evaluateURL(), bytes.NewBuffer(payload))
	if err != nil {
		return err, 0
	}

	var res = &score_api.ScoreRes{}

	err = t.forwardTo(req, res)
	if err != nil {
		return err, 0
	}

	if res.Status == -1 {
		err = errors.New(res.Msg)
		score = -1
	} else {
		err = nil
		score = res.Metrics.Acc
	}

	return err, score
}

func (t CalculateEvaluate) forwardTo(req *http.Request, jsonResp interface{}) (err error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	_, err = t.cli.ForwardTo(req, jsonResp)
	return
}

type CalculateEvaluate struct {
	endpoint string
	cli      utils.HttpClient
}
