package data

import (
	"github.com/PietroCarrara/aenianos"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"log"
)

const MainSession = "session"

var Store = sessions.NewCookieStore([]byte("secret"))

var Db *gorm.DB

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
