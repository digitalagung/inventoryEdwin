package user

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/edwinlab/inventory/mysql"
)


type User struct {
	Id        int64  `db:"id" json:"id"`
	Name 	  string `db:"name" json:"name"`
}

var dbmap = mysql.InitDb()

func GetUsers(c echo.Context) error {
	var users []User
	_, err := dbmap.Select(&users, "SELECT * FROM users")
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return err
	}
	insert, err := dbmap.Exec(`INSERT INTO users (name) VALUES (?)`, user.Name);
	if err != nil {
		return err
	}

	user_id, _ := insert.LastInsertId()
	content := &User{
		Id:   user_id,
		Name: user.Name,
	}
	return c.JSON(http.StatusCreated, content)
}

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user User
	err := dbmap.SelectOne(&user, "SELECT * FROM users WHERE id=? LIMIT 1", id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	err := dbmap.SelectOne(&user, "SELECT * FROM users WHERE id=? LIMIT 1", id)
	if err != nil {
		return err
	}

	_, err = dbmap.Update(&user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user User
	err := dbmap.SelectOne(&user, "SELECT * FROM users WHERE id=? LIMIT 1", id)
	if err != nil {
		return err
	}

	_, err = dbmap.Delete(&user)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}