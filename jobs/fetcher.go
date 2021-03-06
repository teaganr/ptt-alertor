package jobs

import (
	"sync"

	log "github.com/meifamily/logrus"

	"time"

	board "github.com/meifamily/ptt-alertor/models/ptt/board/redis"
)

type Fetcher struct {
}

func NewFetcher() *Fetcher {
	f := new(Fetcher)
	return f
}

func (f Fetcher) Run() {
	boards := new(board.Board).All()

	var wg sync.WaitGroup
	for _, bd := range boards {
		wg.Add(1)
		go func(bd board.Board) {
			defer wg.Done()
			bd.Articles = bd.FetchArticles()
			bd.Save()
			log.WithField("board", bd.Name).Info("Fetched")
		}(*bd)
		time.Sleep(50 * time.Millisecond)
	}
	wg.Wait()
	log.Info("All fetcher done")
}
