package models

import (
	"database/sql"
	"fmt"

	"github.com/alghiffaryfa19/echo-rest/db"
	"github.com/alghiffaryfa19/echo-rest/helpers"
)

type Users struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role string `json:"role"`
}


func CheckLogin(email string, password string) (bool, error) {
	var obj Users
	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM users WHERE email = ?"

	err := con.QueryRow(sqlStatement, email).Scan(
		&obj.Id, &obj.Name, &obj.Email, &pwd, &obj.Role,
	)

	if err == sql.ErrNoRows {
		fmt.Println("Email not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query Error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("Hash and Password doesn't match")
		return false, err
	}

	return true, nil
}