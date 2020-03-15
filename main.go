package main

import (
	"ClassWork_MGO/server/controller"
	"ClassWork_MGO/server/db_mgo"
)

func main() {
	 controller.StartServ()
	defer db_mgo.SessionGlob.Close()

}

