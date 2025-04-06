package main

import (
    "github.com/gin-gonic/gin"
    "go-webservice/config"
    "go-webservice/router"
    "os"
)

func main() {
    dsn := os.Getenv("DB_DSN")
    config.ConnectDatabase(dsn)

    r := router.SetupRouter()
    r.Run(":8080")
}