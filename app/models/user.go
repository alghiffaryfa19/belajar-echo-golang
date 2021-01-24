package models

import (
	"net/http"

	"github.com/alghiffaryfa19/echo-rest/db"
	validator "github.com/go-playground/validator/v10"
)

type RegisterUser struct {
	Id int `json:"id"`
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role string `json:"role"`
}

func Register(name string, email string, password string, role int) (Response, error) {
	var res Response
	v := validator.New()
	peg := RegisterUser{
		Name : name,
		Email : email,
		Password : password,
	}

	err := v.Struct(peg)
	if err != nil {
		return res, err
	}
	con := db.CreateCon()
	sqlStatement := "INSERT users (name, email, password, role) VALUES (?, ?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(name, email, password, role)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"las_inserted_id": lastInsertedId,
	}

	return res, nil
}