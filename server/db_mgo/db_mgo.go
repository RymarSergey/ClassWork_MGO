package db_mgo

import (
	"ClassWork_MGO/server/model"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

var SessionGlob *mgo.Session

func ConnectToDb() {
	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		panic(err)
	}
	SessionGlob = session
	fmt.Println("SessionGlob = ",SessionGlob)

}

func SaveToDb(human *model.Human)  bool{
	if SessionGlob==nil{
		ConnectToDb()
	}
	// добавляем один объект
	err := SessionGlob.DB("humandb").C("humans").Insert(human)
	if err != nil{
		fmt.Println(err)
		return false
	}
	return true
}
func GetFromDb(fname string) []model.Human{
	if SessionGlob==nil{
		ConnectToDb()
	}
	// критерий выборки
	query := bson.M{
		"fname" : fname,
	}
	// объект для сохранения результата
	humans :=[] model.Human{}
	SessionGlob.DB("humandb").C("humans").Find(query).All(&humans)
	return humans
}

func SetToDb(human *model.Human) bool {
	if SessionGlob==nil{
		ConnectToDb()
	}
	err := SessionGlob.DB("humandb").C("humans").Update(bson.M{"fname": human.Fname}, bson.M{"$set":bson.M{"age":human.Age}})
	if err != nil{
		return false
	}
	return true
}
func DeleteHuman(fname string) bool {
	if SessionGlob==nil{
		ConnectToDb()
	}
	// критерий выборки
	query := bson.M{
		"fname" : fname,
	}
	err := SessionGlob.DB("humandb").C("humans").Remove(query)
	if err != nil{
		fmt.Println(err)
		return false
	}
	return true
}

