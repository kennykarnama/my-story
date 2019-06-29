package story

import (
	"context"
	"encoding/json"
	"errors"
	"os"

	"github.com/google/logger"
	"github.com/kennykarnama/my-story/config"
)

const (
	//RepoErrFileNotExist is an error that happens
	//when the file cant be found
	RepoErrFileNotExist = "File Specified Not Found"
)

//Repository is an interface to provide Data Interaction
type Repository interface {
	//GetStories get story struct from json file
	GetStories(ctx context.Context) ([]Story, error)
}

type storyRepository struct {
	config config.Config
}

//NewStoryRepository constructs new repository
func NewStoryRepository(config config.Config) Repository {
	return &storyRepository{config}
}

func (sr *storyRepository) GetStories(ctx context.Context) ([]Story, error) {
	filePath := sr.config.StoryPath
	if filePath == "" {
		logger.Errorf("File specified not found: %s", filePath)
		return nil, errors.New(RepoErrFileNotExist)
	}
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		logger.Errorf("Error opening file: %v", err.Error())
	}
	decoder := json.NewDecoder(file)
	var story []Story
	err = decoder.Decode(&story)
	if err != nil {
		logger.Errorf("Error when unmarshal story: %v", err)
	}
	return story, nil
}
