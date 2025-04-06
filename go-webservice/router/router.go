package router

import (
    "github.com/gin-gonic/gin"
    "go-webservice/controller"
    "go-webservice/middleware"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.Use(middleware.AuthMiddleware())

    api := r.Group("/api")
    {
        api.GET("/courses", controller.GetCourses)
        api.GET("/courses/:id", controller.GetCourse)
    }

    return r
}