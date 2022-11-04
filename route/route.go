package route

import (
	"github.com/gin-gonic/gin"
	"github.com/qinsheng99/go-py/controller"
	"github.com/qinsheng99/go-py/infrastructure/score"
	"net/http"
	"os"
)

func SetRoute(r *gin.Engine) {
	group := r.Group("/v1")
	group.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "success")
	})

	controller.AddRouteScore(
		group, score.NewScore(os.Getenv("EVALUATE"), os.Getenv("CALCULATE")),
	)
}
