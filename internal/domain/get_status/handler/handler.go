package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type getStatusHandler struct {
}

func NewGetStatusHandler() *getStatusHandler {
	return &getStatusHandler{}
}

// @Summary Service Status
// @Description Indicate the service has started up correctly and is ready to accept requests.
// @ID getStatus
// @Success 200 "OK"
// @Router /status [get]
func (h *getStatusHandler) Handle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}
