package controller

import (
	"github.com/B6310851/team05/entity"

	"net/http"

	"github.com/gin-gonic/gin"
)


func CreateQualification(c *gin.Context) {
	var qualification entity.Qualification
	if err := c.ShouldBindJSON(&qualification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&qualification).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": qualification})
}

func GetQualification(c *gin.Context) {
	var qualification entity.Qualification
	id := c.Param("qualification_id")
	if err := entity.DB().Raw("SELECT * FROM qualifications WHERE qualification_id = ?", id).Find(&qualification).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": qualification})
}


func ListQualifications(c *gin.Context) {
	var qualifications []entity.Qualification
	if err := entity.DB().Raw("SELECT * FROM qualifications").Find(&qualifications).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": qualifications})
}
