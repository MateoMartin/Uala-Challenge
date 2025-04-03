package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"uala-challenge/internal/model"
)

//go:generate mockery --name=getTimelineUseCase --structname=GetTimelineUseCase --output=./mocks
type getTimelineUseCase interface {
	GetTimeline(ctx context.Context, userID string) (*model.Timeline, error)
}

type getTimelineHandler struct {
	useCase getTimelineUseCase
}

func NewGetTimelineHandler(useCase getTimelineUseCase) *getTimelineHandler {
	return &getTimelineHandler{
		useCase: useCase,
	}
}

// @Summary Get Timeline
// @Description Get timeline of tweets for a user
// @ID getTimeline
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 "OK"
// @Failure 500 "Internal Server Error"
// @Router /timeline/{user_id} [get]
func (h *getTimelineHandler) Handle(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	timeline, err := h.useCase.GetTimeline(ctx, userID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, timeline)
}
