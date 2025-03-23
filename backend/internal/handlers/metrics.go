package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetMetrics(c *gin.Context) {
    // fake metrics to test
    c.JSON(http.StatusOK, gin.H{
        "cpu_usage":  20.5,
        "ram_usage":  45.3,
        "disk_usage": 60.2,
    })
}
