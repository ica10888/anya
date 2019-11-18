package models

import (
	"github.com/astaxie/beego"
	"github.com/boltdb/bolt"
	"log"
	"os"
	"path"
)

func Init() {
	dir := beego.AppConfig.String("filepath")
	createDirIfNotExist(dir)

	path := path.Join(dir, "bolt.db")
	createFile, err := os.Open(path)
	defer func() {
		createFile.Close()
	}()
	if err != nil && os.IsNotExist(err) {
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
