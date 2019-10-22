package models

import (
	"github.com/astaxie/beego"
	"log"
	"os"
	"path"

	"github.com/boltdb/bolt"
)

func Init() {
	dir := beego.AppConfig.String("filepath")
	_, err := os.Stat(dir)
	if err != nil {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	path := path.Join(dir,"bolt.db")
	createFile,err:=os.Open(path)
	defer func(){
		createFile.Close()
	}()
	if err != nil && os.IsNotExist(err){
		_, err = os.Create(path)
		if err != nil {
			panic(err)
		}
	}

	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}