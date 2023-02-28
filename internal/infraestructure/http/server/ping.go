package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingController struct{}

func (r *PingController) GetPing(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
