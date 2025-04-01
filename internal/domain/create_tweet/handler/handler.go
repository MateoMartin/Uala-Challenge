package handler

import (
	"context"
	"net/http"
	"uala-challenge/internal/model"

	"github.com/gin-gonic/gin"
)

//go:generate mockery --name=createTweetUseCase --structname=CreateTweetUseCase --output=./mocks
type createTweetUseCase interface {
	CreateTweet(ctx context.Context, tweet *model.Tweet) error
}

type createTweetHandler struct {
	useCase createTweetUseCase
}

func NewCreateTweetHandler(useCase createTweetUseCase) *createTweetHandler {
	return &createTweetHandler{
		useCase: useCase,
	}
}

// @Summary Create Tweet
// @Description tweet to post
// @ID createTweet
// @Accept json
// @Produce json
// @Param tweet body createTweetDTO true "Tweet Request"
// @Success 201 "Created"
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /tweets [post]
func (h *createTweetHandler) Handle(ctx *gin.Context) {
	var request createTweetDTO
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	tweet, err := request.toTweet()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.useCase.CreateTweet(ctx, tweet)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, tweet)
}
