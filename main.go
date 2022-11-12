package main

import (
    "github.com/gin-gonic/gin"
    "github.com/hollo08/raftexample/handler"
)

func main() {
    r := gin.Default()
    r.GET("/set", handler.Set)
    r.GET("/get", handler.Get)
    r.GET("/join", handler.Join)
    r.Run(":8989")
}
