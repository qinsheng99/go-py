package sdk

import (
	"bytes"
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

func (t CalculateEvaluate) Evaluate(opt *Calculate) error {
	payload, err := utils.JsonMarshal(opt)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, t.evaluateURL(), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	return t.forwardTo(req, nil)
}

func (t CalculateEvaluate) forwardTo(req *http.Request, jsonResp interface{}) (err error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if jsonResp != nil {
		v := struct {
			Data interface{} `json:"data"`
		}{jsonResp}

		_, err = t.cli.ForwardTo(req, &v)
	} else {
		_, err = t.cli.ForwardTo(req, jsonResp)
	}

	return
}

type CalculateEvaluate struct {
	endpoint string
	cli      utils.HttpClient
}
