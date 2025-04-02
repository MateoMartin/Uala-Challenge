package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

//go:generate mockery --name=followUserUseCase --structname=FollowUserUseCase --output=./mocks
type followUserUseCase interface {
	FollowUser(ctx context.Context, userID string, userIDToFollow string) error
}

type followUserHandler struct {
	useCase followUserUseCase
}

func NewFollowUserHandler(useCase followUserUseCase) *followUserHandler {
	return &followUserHandler{
		useCase: useCase,
	}
}

// @Summary Follow User
// @Description follows an user
// @ID followUser
// @Accept json
// @Produce json
// @Param tweet body followUserDTO true "Follow User"
// @Success 202 "Accepted"
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /follow [post]
func (h *followUserHandler) Handle(ctx *gin.Context) {
	var request followUserDTO
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = h.useCase.FollowUser(ctx, request.UserID, request.UserIDToFollow)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusAccepted)
}
