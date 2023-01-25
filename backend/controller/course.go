package controller

import (
	"github.com/B6310851/team05/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /users

type extendedCourse struct {
	entity.Course
	Qualification_Name string
}

func CreateCourse(c *gin.Context) {

	var course entity.Course
	var qualification entity.Qualification
	var major entity.Major

	if err := c.ShouldBindJSON(&course); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	if tx := entity.DB().Where("qualification_id = ?", course.Qualification_ID).First(&qualification); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "qualification not found"})
		return
	}
	if tx := entity.DB().Where("major_id = ?", course.Major_ID).First(&major); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "major not found"})
		return
	}

	new_course := entity.Course{
		Course_ID:        course.Course_ID,
		Course_Name:      course.Course_Name,
		Datetime:         course.Datetime,
		Major_ID:         course.Major_ID,
		Qualification_ID: course.Qualification_ID,
	}

	if err := entity.DB().Create(&course).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": new_course})

}

// GET /user/:id

func GetCourse(c *gin.Context) {

	var user entity.Course

	id := c.Param("course_id")

	if err := entity.DB().Raw("SELECT * FROM courses WHERE course_id = ?", id).Scan(&user).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}

// GET /users

func ListCourses(c *gin.Context) {

	var courses []extendedCourse

	if err := entity.DB().Preload("Qualification").Preload("Major").Preload("Course").Raw("SELECT c.* , q.qualification_name FROM qualifications q JOIN courses c ON q.qualification_ID = c.qualification_ID").Scan(&courses).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": courses})

}

func DeleteCourse(c *gin.Context) {
	course_id := c.Param("course_id")

	if tx := entity.DB().Exec("DELETE FROM courses WHERE course_id = ?", course_id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "course not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": course_id})
}
