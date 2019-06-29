package story

import (
	"context"

	"github.com/google/logger"
)

//Service provides parser of story
type Service interface {
	//ParseStory provides access to story
	ParseStory(cxt context.Context) ([]Story, error)
}

type storyService struct {
	storyRepo Repository
}

//NewStoryService constructs new story service
func NewStoryService(storyRepo Repository) Service {
	s := &storyService{storyRepo}
	return s
}

func (ss *storyService) ParseStory(ctx context.Context) ([]Story, error) {
	story, err := ss.storyRepo.GetStories(ctx)
	if err != nil {
		logger.Errorf("Error ParseStory: %v", err.Error())
		return nil, err
	}
	return story, nil
}
