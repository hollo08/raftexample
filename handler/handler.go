package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/raft"
	"github.com/hollo08/raftexample/cache"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

const (
	ENABLE_WRITE_TRUE  = int32(1)
	ENABLE_WRITE_FALSE = int32(0)
)

type Handler struct {
	Ctx         *cache.RaftCachedContext
	Log         *log.Logger
	Mux         *http.ServeMux
	EnableWrite int32
}

func (h *Handler) checkWritePermission() bool {
	return atomic.LoadInt32(&h.EnableWrite) == ENABLE_WRITE_TRUE
}

func (h *Handler) Set(c *gin.Context) {
	key, _ := c.GetQuery("key")
	value, _ := c.GetQuery("value")
	if key == "" || value == "" {
		c.String(200, fmt.Sprintf("fail"))
		return
	}
	event := cache.LogEntryData{Key: key, Value: value}
	eventBytes, err := json.Marshal(event)
	if err != nil {
		h.Log.Printf("json.Marshal failed, err:%v", err)
		fmt.Fprint(c.Writer, "internal error\n")
		return
	}

	applyFuture := h.Ctx.RC.Raft.Raft.Apply(eventBytes, 5*time.Second)
	if err := applyFuture.Error(); err != nil {
		h.Log.Printf("raft.Apply failed:%v", err)
		fmt.Fprint(c.Writer, "internal error\n")
		return
	}
	c.String(200, fmt.Sprintf("hello %s, %s\n", key, value))
}

func (h *Handler) Get(c *gin.Context) {
	key, _ := c.GetQuery("key")
	if key == "" {
		c.String(200, fmt.Sprintf("fail"))
		return
	}
	ret := h.Ctx.RC.Cm.Get(key)
	c.String(200, fmt.Sprintf("hello %s\n", ret))
}

func (h *Handler) Join(c *gin.Context) {
	node, _ := c.GetQuery("node")
	if node == "" {
		c.String(200, fmt.Sprintf("fail"))
		return
	}
	addPeerFuture := h.Ctx.RC.Raft.Raft.AddVoter(raft.ServerID(node), raft.ServerAddress(node), 0, 0)
	if err := addPeerFuture.Error(); err != nil {
		h.Log.Printf("Error joining peer to raft, peerAddress:%s, err:%v, code:%d", node, err, http.StatusInternalServerError)
		fmt.Fprint(c.Writer, "internal error\n")
		return
	}
	c.String(200, "ok")
}
