package jike

import (
	"context"
	"net/http"
)

type TopicService struct {
	jike   *Jike
	client *http.Client
	debug  bool
}

func NewTopicService(jike *Jike, c *http.Client, debug bool) *TopicService {
	return &TopicService{jike, c, debug}
}

func (u *TopicService) HotFeed(ctx context.Context, topicId string, limit int, loadMoreKey HotFeedLoadMoreKey) (*HotFeedOutput, error) {
	input := map[string]interface{}{
		"topicId":     topicId,
		"limit":       limit,
		"loadMoreKey": loadMoreKey,
	}
	output := &HotFeedOutput{}
	req := &request{
		Debug:      u.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `/1.0/topics/tabs/selected/feed`,
		Input:      input,
		Output:     &output,
		Context:    ctx,
	}
	return output, req.send(u.jike)
}
