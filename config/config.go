package config

import (
	"github.com/kelseyhightower/envconfig"
)

//Config represents configurable option for this app
type Config struct {
	//StoryPath is where the json file of the story
	//located
	StoryPath string `env:"STORY_PATH" default:"story.json"`
	//StorySongPath is where the storyson mp3 is located
	StorySongPath string `env:"STORY_SONG_PATH" default:"story_song.mp3"`
}

// Get to get defined configuration
func Get() Config {
	cfg := Config{}
	envconfig.MustProcess("", &cfg)
	return cfg
}
