package users

import (
	"log"

	"github.com/branLiang/gin-playground/db"
	"github.com/gin-gonic/gin"
)

func RouterRegister(r *gin.RouterGroup) {
	r.GET("/", ListUsers)
	r.GET("/ids", ListUserIds)
}

// ListUsers return users
func ListUsers(c *gin.Context) {
	users := make([]User, 0)
	rows, err := db.DB.Query("SELECT id, name FROM users LIMIT 10")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Name)
		users = append(users, user)
	}

	c.JSON(200, gin.H{"data": users})
}

func ListUserIds(c *gin.Context) {
	users := make([]UserSimpleSerializer, 0)
	rows, err := db.DB.Query("SELECT id FROM users LIMIT 10")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var user UserSimpleSerializer
		rows.Scan(&user.Id)
		users = append(users, user)
	}

	c.JSON(200, users)
}
