package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/google/logger"

	"github.com/kennykarnama/tview"

	"github.com/kennykarnama/my-story/story"

	"github.com/kennykarnama/my-story/config"
	"github.com/kennykarnama/my-story/song"
)

func initRepo(cfg config.Config) story.Repository {
	repo := story.NewStoryRepository(cfg)
	return repo
}

func initService(repo story.Repository) story.Service {
	svc := story.NewStoryService(repo)
	return svc
}

func initMusicHandler(cfg config.Config) song.Handler {
	musicHandler := song.NewHandler(cfg)
	return musicHandler
}

const logPath = "app.log"

var verbose = flag.Bool("verbose", false, "print info level logs to stdout")

func main() {

	cfg := config.Get()

	repo := initRepo(cfg)
	svc := initService(repo)
	musicHandler := initMusicHandler(cfg)

	app := tview.NewApplication()
	pages := tview.NewPages()
	stories, err := svc.ParseStory(context.Background())
	if err != nil {
		log.Fatalf(err.Error())
	}

	stop := func() {
		musicHandler.Stop()
		app.Stop()
	}

	app.HandleStop(stop)
	pageCount := len(stories)
	for page := 0; page < pageCount; page++ {
		item := stories[page]
		logger.Infof("%v", item)
		func(page int, i story.Story, mh song.Handler) {
			pages.AddPage(fmt.Sprintf("%s", i.Title),
				tview.NewModal().
					SetText(fmt.Sprintf("%s", i.Content)).
					AddButtons([]string{"Next", "Quit"}).
					SetDoneFunc(func(buttonIndex int, buttonLabel string) {
						if buttonIndex == 0 {
							pages.SwitchToPage(fmt.Sprintf("%s", stories[(page+1)%pageCount].Title))
						} else {
							musicHandler.Stop()
							app.Stop()
						}
					}),
				false,
				page == 0)
		}(page, item, musicHandler)
	}

	//done 2 := make(chan bool)
	// c := make(chan os.Signal)
	// signal.Notify(c, os.Interrupt)

	musicHandler.Play()

	if err := app.SetRoot(pages, true).SetFocus(pages).Run(); err != nil {
		panic(err)
	}

}
