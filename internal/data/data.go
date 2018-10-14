package data

import (
	"log"

	"github.com/PietroCarrara/aenianos"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

var r *mux.Router

func init() {
	var err error

	user := "aenianos"
	password := "password"
	dbname := "aenianos"
	Db, err = gorm.Open("mysql", user+":"+password+"@/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	Db.AutoMigrate(&aenianos.User{})
	Db.AutoMigrate(&aenianos.Genero{})
}

func GetRouter() *mux.Router {
	return r
}

func SetRouter(router *mux.Router) {
	r = router
}
