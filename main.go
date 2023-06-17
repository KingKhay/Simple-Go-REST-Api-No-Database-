package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Dob       string `json:"dob"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

var users = []User{
	{
		ID:        1,
		FirstName: "King",
		LastName:  "Khay",
		Dob:       "1990-06-12",
		Email:     "kingkhay@gmail.com",
		Username:  "khay",
		Password:  "01234566778",
	},
	{
		ID:        2,
		FirstName: "King",
		LastName:  "Khay",
		Dob:       "1990-06-12",
		Email:     "kingkhay@gmail.com",
		Username:  "khay",
		Password:  "01234566778",
	},
	{
		ID:        3,
		FirstName: "King",
		LastName:  "Khay",
		Dob:       "1990-06-12",
		Email:     "kingkhay@gmail.com",
		Username:  "khay",
		Password:  "01234566778",
	},
	{
		ID:        4,
		FirstName: "King",
		LastName:  "Khay",
		Dob:       "1990-06-12",
		Email:     "kingkhay@gmail.com",
		Username:  "khay",
		Password:  "01234566778",
	},
	{
		ID:        5,
		FirstName: "King",
		LastName:  "Khay",
		Dob:       "1990-06-12",
		Email:     "kingkhay@gmail.com",
		Username:  "khay",
		Password:  "01234566778",
	},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func addUser(c *gin.Context) {
	var newUser User
	//BindJSON data to the newly created User "newUser"
	if err := c.BindJSON(&newUser); err != nil {
		log.Fatal(err)
		return
	}
	//Add the new User to the Slice
	users = append(users, newUser)

	//Return JSON format of newly created User
	c.IndentedJSON(http.StatusCreated, newUser)
}

func getUserById(c *gin.Context) {
	//Get the user ID from the request parameters
	userIDStr := c.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user Id"})
		return
	}
	//Find the user by ID
	var foundUser *User
	for _, user := range users {
		if user.ID == userID {
			foundUser = &user
			break
		}
	}
	//Check if user was found
	if foundUser == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	//Return the user as JSON response
	c.JSON(http.StatusOK, foundUser)
}

func updateUser(c *gin.Context) {
	var updatedUser User
	err := c.ShouldBindJSON(&updatedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	userIdStr := c.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
	}

	for i, user := range users {
		if user.ID == userId {
			users[i].FirstName = updatedUser.FirstName
			users[i].LastName = updatedUser.LastName
			users[i].Dob = updatedUser.Dob
			users[i].Email = updatedUser.Email
			users[i].Password = updatedUser.Password
			users[i].Username = updatedUser.Username
			break
		}
	}

	c.IndentedJSON(http.StatusOK, &updatedUser)
}

func deleteUser(c *gin.Context) {
	userIdStr := c.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
		return
	}
	for i, user := range users {
		if user.ID == userId {
			// Remove the user from the slice
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "User deleted successfully",
			})
			return
		}
	}

	// If user not found, return an error response
	c.JSON(http.StatusNotFound, gin.H{
		"error": "User not found",
	})
}

func main() {

	router := gin.Default()

	router.GET("/users", getUsers)

	router.GET("/users/:id", getUserById)

	router.POST("/users", addUser)

	router.PUT("/users/:id", updateUser)

	router.DELETE("/users/:id", deleteUser)

	router.Run("localhost:9000")
}
