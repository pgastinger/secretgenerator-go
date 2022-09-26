package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type response struct {
	UUID string `json:"uuid"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", getUUID)
	router.Run("0.0.0.0:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getUUID(c *gin.Context) {
	id := uuid.New()
	c.IndentedJSON(http.StatusOK, &response{UUID: id.String()})
}
