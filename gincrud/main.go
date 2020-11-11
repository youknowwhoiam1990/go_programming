package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db, err = sql.Open("mysql", "root:hello1234@tcp(127.0.0.1:3306)/abhi")

type User struct {
	Id          int    `db:"id"`
	Username    string `db:"username" form:"username"`
	FirstName   string `db:"first_name" form:"first_name"`
	LastName    string `db:"last_name" form:"last_name"`
	Email       string `db:"email" form:"email"`
	MobilePhone string `db:"mobile_phone" form:"mobile_phone"`
}

func getAll(c *gin.Context) {
	var (
		user  User
		users []User
	)
	//This returns a row
	rows, err := db.Query("select id, username, first_name, last_name, email, mobile_phone from user;")

	if err != nil {
		fmt.Print(err.Error())
	}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.MobilePhone)
		users = append(users, user)
	}

	defer rows.Close()
	c.JSON(http.StatusOK, users)
}

func xyz(c *gin.Context) {
	Username := c.PostForm("username")
	FirstName := c.PostForm("first_name")
	LastName := c.PostForm("last_name")
	Email := c.PostForm("email")
	MobilePhone := c.PostForm("mobile_phone")

	if Username == "" || FirstName == "" || Email == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Please fill all mandatory field"),
		})
		return
	}

	if isEmailOrUsernameAlreadyExist(Username, Email) {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Username or email is already in use"),
		})
		return
	}

	stmt, err := db.Prepare("insert into user (username,first_name, last_name, email, mobile_phone) values(?,?,?,?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}

	_, err = stmt.Exec(Username, FirstName, LastName, Email, MobilePhone)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("successfully created"),
	})
}

func isEmailOrUsernameAlreadyExist(username, email string) bool {
	var Id int
	err := db.QueryRow("SELECT id FROM user WHERE username= ? OR email= ?", username, email).Scan(&Id)

	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that ID.")
		return false
	case err != nil:
		log.Fatal(err)
		return false
	default:
		log.Printf("Found")
		return true
	}
}

func update(c *gin.Context) {
	Id, err := strconv.Atoi(c.Param("id")) //strconv.Atoi converts string to int
	Username := c.PostForm("username")
	FirstName := c.PostForm("first_name")
	LastName := c.PostForm("last_name")
	Email := c.PostForm("email")
	MobilePhone := c.PostForm("mobile_phone")

	//checking if anything field is vacant
	if Username == "" || FirstName == "" || Email == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Please fill all mandatory field"),
		})
		return
	}

	stmt, err := db.Prepare("update user set username= ?,first_name= ?, last_name= ?, email= ?, mobile_phone= ? where id= ?;")
	if err != nil {
		fmt.Print(err.Error())
	}

	_, err = stmt.Exec(Username, FirstName, LastName, Email, MobilePhone, Id)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer stmt.Close()

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully updated"),
	})
}

func getById(c *gin.Context) {
	var user User

	id := c.Param("id")
	row := db.QueryRow("select id, username, first_name, last_name, email, mobile_phone from user where id = ?;", id)

	err = row.Scan(&user.Id, &user.Username, &user.FirstName, &user.LastName, &user.Email, &user.MobilePhone)
	if err != nil {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func delete(c *gin.Context) {
	id := c.Param("id")
	stmt, err := db.Prepare("delete from user where id= ?;")

	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(id) //check

	if err != nil {
		fmt.Print(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted user with ID : %s", id),
	})
}

func createTable() {
	stmt, err := db.Prepare("CREATE TABLE user (id int NOT NULL AUTO_INCREMENT, username varchar(40), first_name varchar(40), last_name varchar(40), email varchar(60), mobile_phone varchar(15), PRIMARY KEY (id));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec() //check
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Table is successfully created....")
	}
}

func main() {
	createTable() //check

	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close() //check

	err = db.Ping() //check
	if err != nil {
		fmt.Print(err.Error())
	}

	router := gin.Default()
	router.GET("/api/user/:id", getById)
	router.GET("/api/users", getAll)
	router.POST("/api/xyz", xyz)
	router.PUT("/api/user/:id", update)
	router.DELETE("/api/user/:id", delete)
	router.Run(":8000")
}
