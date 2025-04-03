package event

import "uala-challenge/internal/model"

type inMemoryEventPublisher struct {
	eventChan chan<- *model.TweetCreatedEvent
}

func NewInMemoryEventPublisher(ch chan<- *model.TweetCreatedEvent) *inMemoryEventPublisher {
	return &inMemoryEventPublisher{eventChan: ch}
}

func (p *inMemoryEventPublisher) Publish(event *model.TweetCreatedEvent) error {
	go func() {
		p.eventChan <- event
	}()
	return nil
}
