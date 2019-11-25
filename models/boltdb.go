package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/boltdb/bolt"
	"log"
	"os"
	"path"
)

var DataBase *bolt.DB

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

	DataBase, err = bolt.Open(path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	//创建 Bucket
	DataBase.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("TitleBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
