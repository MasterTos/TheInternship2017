package main

import (
   "database/sql"
   "golang.org/x/crypto/bcrypt"
   "github.com/gin-gonic/gin"
   _ "github.com/go-sql-driver/mysql"
	"net/http"
	"fmt"
)

//Account use for parse value
type Account struct {
	ID int `json:"id"`
	Username string `json:"username"`
}

var (
	db *sql.DB
	err error
)


func main() {
	db, err = sql.Open("mysql", "test:test@tcp(127.0.0.1:3306)/gotest")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	router := gin.Default()
	router.POST("/signIn", signIn)
	router.POST("/signUp", signUp)

	router.Run(":8080")
}

func signIn(c *gin.Context) {
	var (
		account Account
		hashedPassword string
		result gin.H
	)
	username := c.PostForm("username")
	password := c.PostForm("password")
	row := db.QueryRow("select id, username, hashedPassword from users where username = ?;", username)
	err = row.Scan(&account.ID, &account.Username, &hashedPassword)
	if err != nil {
		result = gin.H{
			"result" : nil,
			"message" : "Incorrect username",
		}
		fmt.Println(err.Error())
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
		if err != nil {
			result = gin.H{
				"result" : nil,
				"message" : "Incorrect password",
			}
			fmt.Println(err.Error())
		} else {
			fmt.Println(account)
			result = gin.H{
				"result" : account,
				"message" : "Login successfully",
			}
		}
	}

	c.JSON(http.StatusOK, result)

}

func signUp(c *gin.Context) {
	var (
		hashedPassword []byte
		result gin.H
	)
	username := c.PostForm("username")
	password := []byte(c.PostForm("password"))
	stmt, err := db.Prepare("insert into users (username, hashedPassword) value(?,?);")
	if err != nil {
		result = gin.H{
			"result" : false,
			"message" : "something error in database",
		}
		fmt.Println(err.Error())
	}
	hashedPassword, err = bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	_, err = stmt.Exec(username, hashedPassword)
	if err != nil {
		result = gin.H{
			"result" : false,
			"message" : "username is already to used",
		}
		fmt.Println(err.Error())
	} else {
		result = gin.H{
			"result" : true,
			"message" : fmt.Sprintf(" %s successfully created", username),
		}
	}

	c.JSON(http.StatusOK, result)

}