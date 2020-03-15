package controller

import (
	"ClassWork_MGO/server/db_mgo"
	"ClassWork_MGO/server/model"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"net/http"
)

type humansGlobal struct {
	Exporthumans []model.Human
}
var humans_Global humansGlobal
func postingSaveHuman(c *gin.Context) {
	tempHuman := model.Human{
		Fname: c.PostForm("fname"),
		Lname: c.PostForm("lname"),
		Age:   c.PostForm("age"),
		Id:    bson.NewObjectId(),
	}

	if db_mgo.SaveToDb(&tempHuman) {
		c.String(http.StatusCreated, " Human saved true ")
		return
	} else {
		c.String(http.StatusOK, " Human saved false ")
	}

}

func getHuman(c *gin.Context) {
	fname:= c.Query("fname")
	if len(fname)==0 {
		c.String(http.StatusBadRequest, "Get Human fname is  empty string")
		return
	}
	humans_Global.Exporthumans = db_mgo.GetFromDb(fname)
	c.HTML(http.StatusOK,"humanForm.html", &humans_Global)
}

func SetHuman(c *gin.Context) {
	tempHuman := model.Human{
		Fname: c.PostForm("fname"),
		Lname: c.PostForm("lname"),
		Age:   c.PostForm("age"),
		Id:    bson.NewObjectId(),
	}
	if db_mgo.SetToDb(&tempHuman) {
		c.String(http.StatusCreated, " Human vas updated ")
		return
	} else {
		c.String(http.StatusOK, " Human vas't updated ")
	}

}
func deleteHuman(c *gin.Context) {
	fname:= c.Query("fname")
	if len(fname)==0 {
		c.String(http.StatusBadRequest, "Get Human fname is  empty string")
		return
	}
	if db_mgo.DeleteHuman(fname) {
		c.String(http.StatusCreated, " Human deleted true ")
		return
	} else {
		c.String(http.StatusOK, " Human deleted false ")
	}

}