package controller

import (
	"ClassWork_MGO/server/db_mgo"
	"github.com/gin-gonic/gin"
)

func StartServ() {
	db_mgo.ConnectToDb()
	router:=gin.Default()

	router.StaticFile("/choice","static/StartPage.html")
	router.StaticFile("/humans/save", "static/SaveHuman.html")
	router.StaticFile("/humans/get", "static/GetHuman.html")
	router.StaticFile("/humans/update", "static/SetHuman.html")
	router.StaticFile("/humans/delete", "static/DeleteHuman.html")
	router.POST("/bd",postingSaveHuman)
	router.GET("/bd",getHuman)
	router.GET("/bd/delete",deleteHuman)
	router.POST("/bd/put",SetHuman)

	router.LoadHTMLFiles("./static/humanForm.html")
	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
