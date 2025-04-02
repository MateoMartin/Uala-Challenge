package usecase

import (
	"context"
	followuser "uala-challenge/internal/domain/follow_user"
	"uala-challenge/utils/logger"
)

//go:generate mockery --name=userRepository --structname=UserRepository --output=./mocks
type userRepository interface {
	FollowUser(ctx context.Context, userID string, userIDToFollow string) error
}

type followUserUseCase struct {
	userRepository userRepository
}

func NewFollowUserUseCase(userRepository userRepository) *followUserUseCase {
	return &followUserUseCase{
		userRepository: userRepository,
	}
}

func (useCase *followUserUseCase) FollowUser(ctx context.Context, userID string, userIDToFollow string) error {
	err := useCase.userRepository.FollowUser(ctx, userID, userIDToFollow)
	if err != nil {
		logger.GetLogger().Errorf("error following user, user_id_to_follow: %s,  err: %s", userID, userIDToFollow, err.Error())
		return followuser.ErrInternal
	}
	return nil
}
