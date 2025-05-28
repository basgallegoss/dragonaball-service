package router

import (
	"net/http"

	"github.com/basgallegoss/dragonball-service/internal/application"
	"github.com/basgallegoss/dragonball-service/internal/infrastructure/api"
	"github.com/basgallegoss/dragonball-service/internal/infrastructure/db"
	"github.com/gin-gonic/gin"
)

func SetupRouter(dsn, apiURL string) (*gin.Engine, error) {
	repo, err := db.NewPostgresRepository(dsn)
	if err != nil {
		return nil, err
	}
	ext := api.NewDragonballAPI(apiURL)
	svc := application.NewCharacterService(repo, ext)

	r := gin.Default()
	r.POST("/characters", func(c *gin.Context) {
		var req struct {
			Name string `json:"name"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		char, err := svc.GetOrCreate(req.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, char)
	})
	return r, nil
}
