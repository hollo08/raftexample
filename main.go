package main

import (
    "github.com/gin-gonic/gin"
    "github.com/hollo08/raftexample/handler"
    "log"
)

func main() {
    r := gin.Default()
    r.GET("/set", handler.Set)
    r.GET("/get", handler.Get)
    r.GET("/join", handler.Join)
    if err := r.Run(":8989"); err != nil {
        log.Fatalf("server run: %s", err)
    }
}
