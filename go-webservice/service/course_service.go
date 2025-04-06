package service

import (
    "errors"
    "go-webservice/model"
)

var courses = []model.Course{
    {ID: "1", Title: "Go Basics", Description: "Learn Go"},
}

func GetAllCourses() []model.Course {
    return courses
}

func GetCourseByID(id string) (model.Course, error) {
    for _, course := range courses {
        if course.ID == id {
            return course, nil
        }
    }
    return model.Course{}, errors.New("course not found")
}