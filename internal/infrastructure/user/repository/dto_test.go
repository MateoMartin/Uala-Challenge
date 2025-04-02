package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"uala-challenge/internal/model"
)

func TestNewDTOFromModel(t *testing.T) {
	user := &model.User{
		ID:        "dbc8370b-1e7d-4d1b-a131-fb0046ec9929",
		Followers: []string{"73790ac3-5d23-4131-b2f9-4ad1e3da61f6"},
		Following: []string{"6546542-5d23-4131-b2f9-4ad1e3da61f6"},
	}

	dto := newDTOFromModel(user)

	assert.Equal(t, user.ID, dto.ID)
	assert.Equal(t, user.Followers, dto.Followers)
	assert.Equal(t, user.Following, dto.Following)
}

func TestUserDTO_ToModel(t *testing.T) {
	dto := &userDTO{
		ID:        "1",
		Followers: []string{"73790ac3-5d23-4131-b2f9-4ad1e3da61f6"},
		Following: []string{"6546542-5d23-4131-b2f9-4ad1e3da61f6"},
	}

	user := dto.toModel()

	assert.Equal(t, dto.ID, user.ID)
	assert.Equal(t, dto.Followers, user.Followers)
	assert.Equal(t, dto.Following, user.Following)
}
