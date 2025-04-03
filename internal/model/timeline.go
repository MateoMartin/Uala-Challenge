package model

type Timeline struct {
	UserID string   `json:"user_id"`
	Tweets []*Tweet `json:"tweets"`
}

func NewTimeline(userID string) *Timeline {
	return &Timeline{
		UserID: userID,
		Tweets: []*Tweet{},
	}
}

func (t *Timeline) AddTweet(tweet *Tweet) {
	t.Tweets = append(t.Tweets, tweet)
}
