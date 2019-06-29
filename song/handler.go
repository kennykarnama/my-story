package song

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"

	"github.com/google/logger"

	"github.com/kennykarnama/my-story/config"
)

//Handler handles song of the song
type Handler interface {
	//Play the music
	Play()
	//Stop will gracefully stops the music
	Stop()

	Patrol()
}

type handler struct {
	cfg       config.Config
	songFile  *os.File
	musicDone chan bool
	quit      chan bool
	runner    *exec.Cmd
}

//NewHandler is a song handler
func NewHandler(config config.Config) Handler {
	return &handler{
		cfg:       config,
		musicDone: make(chan bool),
		quit:      make(chan bool),
	}
}
func (h *handler) Play() {
	filePath := h.cfg.StorySongPath
	//decode
	h.PlayBackground(filePath)
}

func (h *handler) Patrol() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for {
		select {
		case <-h.quit:
			return
		case <-c:
			h.Stop()
			return
		case <-h.musicDone:
			return
		}
	}
}

//PlayBackground will invoke ffplayer command to
//play the music
func (h *handler) PlayBackground(filepath string) {
	h.runner = exec.Command("ffplay", "-nodisp", fmt.Sprintf("%s", filepath))
	go func(f string, cmd *exec.Cmd) {
		logger.Infof("Executing FFPlayer with the filepath: %s", f)
		_, err := cmd.CombinedOutput()
		if err != nil {
			logger.Errorf("Error playing file mp3 %s: %v", f, err)
			//h.Stop()
		}
		logger.Infof("Music done playing...")
		//logger.Infof(string(out))
		//h.musicDone <- true
		return
	}(filepath, h.runner)
}

func (h *handler) Stop() {
	logger.Infof("Stopping music...")
	err := h.runner.Process.Kill()
	if err != nil {
		logger.Errorf("Failed to kill background process: %v", err.Error())
	}
	//h.quit <- true
}
