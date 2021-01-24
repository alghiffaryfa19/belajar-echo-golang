package main

import (
	"github.com/alghiffaryfa19/echo-rest/db"
	"github.com/alghiffaryfa19/echo-rest/routes"
)

func main(){
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}