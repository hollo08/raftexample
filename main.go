package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hollo08/raftexample/cache"
	"github.com/hollo08/raftexample/handler"
	"log"
	"os"
)

func main() {
	st := &cache.StCached{
		Opts: cache.NewOptions(),
		Log:  log.New(os.Stderr, "stCached: ", log.Ldate|log.Ltime),
		Cm:   cache.NewCacheManager(),
	}
	ctx := &cache.StCachedContext{St: st}

	raft, err := cache.NewRaftNode(st.Opts, ctx)
	if err != nil {
		st.Log.Fatal(fmt.Sprintf("new raft node failed:%v", err))
	}
	st.Raft = raft

	if st.Opts.JoinAddress != "" {
		err = cache.JoinRaftCluster(st.Opts)
		if err != nil {
			st.Log.Fatal(fmt.Sprintf("join raft cluster failed:%v", err))
		}
	}
	logger := log.New(os.Stderr, "httpserver: ", log.Ldate|log.Ltime)
	h := &handler.Handler{
		Ctx:         ctx,
		Log:         logger,
		Mux:         nil,
		EnableWrite: handler.ENABLE_WRITE_FALSE,
	}
	r := gin.Default()
	r.GET("/set", h.Set)
	r.GET("/get", h.Get)
	r.GET("/join", h.Join)
	go func() {
		if err := r.Run(st.Opts.HttpAddress); err != nil {
			log.Fatalf("server run: %s", err)
		}
	}()

	// monitor leadership
	for {
		select {
		case leader := <-st.Raft.LeaderNotifyCh:
			if leader {
				st.Log.Println("become leader, enable write api")
				h.EnableWrite = handler.ENABLE_WRITE_TRUE
			} else {
				st.Log.Println("become follower, close write api")
				h.EnableWrite = handler.ENABLE_WRITE_FALSE
			}
		}
	}
}
