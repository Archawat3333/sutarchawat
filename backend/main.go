package main

import (
	"github.com/B6310851/team05/controller"
	"github.com/B6310851/team05/entity"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT,DELETE")

		if c.Request.Method == "OPTIONS" {

			c.AbortWithStatus(204)

			return

		}

		c.Next()

	}

}

func main() {

	entity.SetupDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())

	// User Routes

	r.GET("/course", controller.ListCourses)
	r.GET("/course/:course_id", controller.GetCourse)
	r.POST("/courses", controller.CreateCourse)
	r.DELETE("/courses/:course_id", controller.DeleteCourse)

	r.GET("/qualifications", controller.ListQualifications)
	r.GET("/qualification/:qualification_id", controller.GetQualification)
	r.POST("qualifications", controller.CreateQualification)
	r.DELETE("/qualifications/:qualification_id", controller.DeleteQualification)
	r.GET("/qualification", controller.ListQualificationName)

	r.GET("/majors", controller.ListMajors)
	r.GET("/major/:major_id", controller.GetMajor)
	r.POST("majors", controller.CreateMajor)
	r.DELETE("/majors/:major_id", controller.DeleteMajor)

	// Run the server

	r.Run()

}
