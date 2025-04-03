package usecase

import (
	"context"
	"uala-challenge/internal/model"
)

//go:generate mockery --name=timelineRepository --structname=TimelineRepository --output=./mocks
type timelineRepository interface {
	GetTimelineByUserID(ctx context.Context, userID string) (*model.Timeline, error)
}

type getTimelineUseCase struct {
	timelineRepository timelineRepository
}

func NewTimelineUseCase(timelineRepository timelineRepository) *getTimelineUseCase {
	return &getTimelineUseCase{
		timelineRepository: timelineRepository,
	}
}

func (useCase *getTimelineUseCase) GetTimeline(ctx context.Context, userID string) (*model.Timeline, error) {
	return useCase.timelineRepository.GetTimelineByUserID(ctx, userID)
}
