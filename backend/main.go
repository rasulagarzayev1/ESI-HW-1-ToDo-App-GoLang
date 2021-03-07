package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/stretchr/testify/assert"
	_ "testing"
)

type Tasks struct {
	ID     int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Title  string `gorm:"not null" form:"title" json:"title"`
	Status bool   `gorm:"not null" form:"status" json:"status"`
	Level  int    `gorm: "default:1" form:"level" json:"level"`
}

// DB Functions

// Create DB
func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./data.db")
	db.LogMode(true) // Error
	if err != nil {
		panic(err)
	} // Creating the table
	if !db.HasTable(&Tasks{}) {
		db.CreateTable(&Tasks{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Tasks{})
	}

	return db
}

func AddTask(c *gin.Context) {

	db := InitDb()
	defer db.Close()

	var task Tasks
	c.Bind(&task)
	fmt.Println("skereeeeeeeeeeeeeeeee")
	fmt.Println(c.Request.Header)
	fmt.Println("skereeeeeeeeeeeeeeeee")


	if task.Title != "" {
		fmt.Print("entro a la creacio	")
		db.Create(&task)
		c.JSON(201, gin.H{"success": task})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

func GetTasks(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	var task []Tasks
	// SELECT * FROM users
	db.Find(&task)

	// Display JSON result
	c.JSON(200, task)
}

func GetTask(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	id := c.Params.ByName("id")
	var task Tasks
	// SELECT * FROM users WHERE id = 1;
	db.First(&task, id)

	if task.ID != 0 {
		// Display JSON result
		c.JSON(200, task)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Task not found"})
	}
}

func UpdateTask(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	// Get id user
	id := c.Params.ByName("id")
	var task Tasks
	// SELECT * FROM users WHERE id = 1;
	db.First(&task, id)

	if task.Title != "" {

		if task.ID != 0 {
			var newTask Tasks
			c.Bind(&newTask)

			result := Tasks{
				ID:     task.ID,
				Title:  newTask.Title,
				Status: newTask.Status,
				Level:  newTask.Level,
			}

			db.Save(&result)
			c.JSON(200, gin.H{"success": result})
		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "Task not found"})
		}

	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}
}

func DeleteTask(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	// Get id user
	id := c.Params.ByName("id")
	var task Tasks
	// SELECT * FROM users WHERE id = 1;
	db.First(&task, id)

	if task.ID != 0 {
		// DELETE FROM users WHERE id = user.Id
		db.Delete(&task)
		// Display JSON result
		c.JSON(200, gin.H{"success": "Task #" + id + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Task not found"})
	}
}

// DB FUnctions

func main() {
	r := setupRouter()
	r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.POST("/tasks", AddTask)
		v1.GET("/tasks", GetTasks)
		v1.GET("/tasks/:id", GetTask)
		v1.PUT("/tasks/:id", UpdateTask)
		v1.DELETE("/tasks/:id", DeleteTask)
	}
	return r
}
