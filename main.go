package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Users []User //nil
// var Users2 = []User{} //empty slice []

func main() {
	r := gin.Default()

	userRoutes := r.Group("/users")
	userRoutes.GET("/", GetUsers)
	userRoutes.POST("/", CreateUser)
	userRoutes.PUT("/:id", EditUser) // PUT /users/355=sdfsfgg=9475
	// localhost/users/

	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}
}

func GetUsers(c *gin.Context) {
	c.JSON(200, Users)
}

func CreateUser(c *gin.Context) {
	var reqBody User
	// ShouldBindJSON accepts a pointer to a structure
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		fmt.Println(reqBody.Name)
		fmt.Println(reqBody.Age)
		return
	}

	reqBody.ID = uuid.New().String()

	Users = append(Users, reqBody)

	c.JSON(200, gin.H{
		"error": false,
	})

}

func EditUser(c *gin.Context) {
	id := c.Param("id")

	var reqBody User
	// ShouldBindJSON accepts a pointer to a structure
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		fmt.Println(reqBody.Name)
		fmt.Println(reqBody.Age)
		return
	}

	for i, u := range Users {
		if u.ID == id {
			Users[i].Name = reqBody.Name
			Users[i].Age = reqBody.Age

			c.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid uuid",
	})
}
