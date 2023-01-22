package controller

import (
	"github.com/B6310851/team05/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /users

func CreateCourse(c *gin.Context) {

	var course entity.Course

	if err := c.ShouldBindJSON(&course); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&course).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": course})

}

// GET /user/:id

func GetCourse(c *gin.Context) {

	var user entity.Course

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM users WHERE course_id = ?", id).Scan(&user).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}

// GET /users

func ListCourses(c *gin.Context) {

	var courses []entity.Course

	if err := entity.DB().Raw("SELECT * FROM courses").Scan(&courses).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": courses})

}
