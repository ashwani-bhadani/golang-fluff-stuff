package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "go-webservice/service"
    "go-webservice/util"
)

func GetCourses(c *gin.Context) {
    c.JSON(http.StatusOK, service.GetAllCourses())
}

func GetCourse(c *gin.Context) {
    id := c.Param("id")
    course, err := service.GetCourseByID(id)
    if err != nil {
        util.HandleError(c, http.StatusNotFound, err.Error())
        return
    }
    c.JSON(http.StatusOK, course)
}