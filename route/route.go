package route

import (
	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-py/controller"
	"github.com/qinsheng99/go-py/infrastructure/score"
	"net/http"
)

func SetRoute(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "success")
	})

	controller.AddRouteScore(
		r,
		score.NewScore(
			"/opt/app/go-py/py/evaluate.py",
			"/opt/app/go-py/py/calculate_fid.py"),
	)
}
