package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/gorm"

	"mycrudapp/database"
)


type Student struct {
	gorm.Model
	Name string `gorm: "size:255"not null`
	Roll_Number int `gorm: "int"not null`
}



func main() {
	loadEnv()
	loadDatabase()
	loadRoutes()	
} 

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&Student{})
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func loadRoutes() {
	r := gin.Default()

	r.GET(
		"/",
		func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello WORLD",
			})
		},
	)

	r.GET("/students", func(c *gin.Context) {
		students := []Student{}
		database.Database.Find(&students)
		c.JSON(200, students)
	})

	r.POST("/student", func(c *gin.Context) {
		student := Student{}
		err := c.ShouldBindJSON(&student)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.BindJSON(&student)

		err_db := database.Database.Create(&student)

		if err_db.Error != nil {
			c.JSON(400, gin.H{"error": err_db.Error})
			return
		}
		
		c.JSON(200, student)
	})
	
	r.Run(":8080")
}