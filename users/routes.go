package users

import (
	"log"

	"github.com/branLiang/gin-playground/db"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func RouterRegister(r *gin.RouterGroup) {
	r.GET("/", ListUsers)
}

// ListUsers return users
func ListUsers(c *gin.Context) {
	users := []User{}
	rows, err := db.DB.Query("SELECT * FROM users LIMIT 10")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Name)
		users = append(users, user)
	}
	usersJson, _ := json.Marshal(&users)
	c.JSON(200, gin.H{
		"data": usersJson,
	})
}
