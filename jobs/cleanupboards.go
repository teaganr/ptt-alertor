package jobs

import (
	log "github.com/meifamily/logrus"
	"github.com/garyburd/redigo/redis"
	"github.com/meifamily/ptt-alertor/connections"
	"github.com/meifamily/ptt-alertor/myutil"
)

type CleanUpBoards struct {
}

func (c *CleanUpBoards) Run() {
	conn := connections.Redis()
	defer conn.Close()
	_, err := redis.Strings(conn.Do("DEL", "boards"))
	boards, err := redis.Strings(conn.Do("KEYS", "board:*"))
	for _, board := range boards {
		_, _ = conn.Do("DEL", board)
	}
	if err != nil {
		log.WithField("runtime", myutil.BasicRuntimeInfo()).WithError(err).Error()
	}
	log.Info("clean up boards")
}
