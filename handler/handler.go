package handler

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func Set(c *gin.Context) {
    key, _ := c.GetQuery("key")
    value, _ := c.GetQuery("value")
    if key == "" || value == "" {
        c.String(200, fmt.Sprintf("fail"))
        return
    }
    c.String(200, fmt.Sprintf("hello %s, %s\n", key, value))
}

func Get(c *gin.Context) {
    key, _ := c.GetQuery("key")
    if key == "" {
        c.String(200, fmt.Sprintf("fail"))
        return
    }
    c.String(200, fmt.Sprintf("hello %s\n", key))
}

func Join(c *gin.Context) {
    node, _ := c.GetQuery("node")
    if node == "" {
        c.String(200, fmt.Sprintf("fail"))
        return
    }
    c.String(200, fmt.Sprintf("hello %s\n", node))
}
