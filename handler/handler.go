package handler

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func Set(c *gin.Context) {
    name := c.DefaultQuery("name", "jack")
    c.String(200, fmt.Sprintf("hello %s\n", name))
}

func Get(c *gin.Context) {
    name := c.DefaultQuery("name", "jack")
    c.String(200, fmt.Sprintf("hello %s\n", name))
}

func Join(c *gin.Context) {
    name := c.DefaultQuery("name", "jack")
    c.String(200, fmt.Sprintf("hello %s\n", name))
}
