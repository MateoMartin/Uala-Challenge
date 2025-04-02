package repository

import (
	"context"
	"sync"
	"uala-challenge/internal/model"
)

type inMemoryUserRepository struct {
	usersByID map[string]*userDTO
	mutex     sync.RWMutex
}

func NewInMemoryUserRepository() *inMemoryUserRepository {
	usersByID := make(map[string]*userDTO, 0)

	return &inMemoryUserRepository{
		usersByID: usersByID,
	}
}

func (r *inMemoryUserRepository) FollowUser(ctx context.Context, userID string, UserIDToFollow string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	userDto := r.usersByID[userID]
	userToFollowDto := r.usersByID[UserIDToFollow]
	if userDto == nil {
		userDto = newDTOFromModel(model.NewUser())
	}
	if userToFollowDto == nil {
		userToFollowDto = newDTOFromModel(model.NewUser())
	}

	user := userDto.toModel()
	userToFollow := userToFollowDto.toModel()

	user.AddFollowing(UserIDToFollow)
	r.usersByID[userID] = newDTOFromModel(user)

	userToFollow.AddFollower(userID)
	r.usersByID[UserIDToFollow] = newDTOFromModel(userToFollow)

	return nil
}

func (r *inMemoryUserRepository) GetUser(ctx context.Context, userID string) (*model.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	userDto := r.usersByID[userID]

	return userDto.toModel(), nil
}
