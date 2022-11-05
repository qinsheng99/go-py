package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/qinsheng99/go-py/api/score_api"
	"github.com/qinsheng99/go-py/app"
	"github.com/qinsheng99/go-py/domain/score"
	"log"
	"net/http"
)

func AddRouteScore(r *gin.RouterGroup, s score.Score) {
	baseScore := BaseScore{
		s: app.NewScoreService(s),
	}

	func() {
		r.POST("/evaluate", baseScore.Evaluate)
		r.POST("/calculate", baseScore.Calculate)
	}()

}

type BaseScore struct {
	s app.ScoreService
}

func (b *BaseScore) Evaluate(c *gin.Context) {
	col := score_api.Score{}
	if err := c.ShouldBindBodyWith(&col, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var res score_api.ScoreRes
	err := b.s.Evaluate(col, &res)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (b *BaseScore) Calculate(c *gin.Context) {
	col := score_api.Score{}
	if err := c.ShouldBindBodyWith(&col, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	go func() {
		var res score_api.ScoreRes
		err := b.s.Calculate(col, &res)
		if err != nil {
			log.Printf("err is %v", err)
			//c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}

		var reserr error
		var data float64

		if res.Status == -1 {
			reserr = errors.New(res.Msg)
			data = -1
		} else {
			reserr = nil
			data = res.Data
		}

		log.Printf("user_name is %s", col.UserName)
		log.Printf("res is %v", data)
		if reserr != nil {
			log.Printf("err is %v", reserr.Error())
		}
	}()

	c.JSON(http.StatusOK, "success")
}
