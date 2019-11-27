package models

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/boltdb/bolt"
	"log"
	"strings"
	"time"
)

var (
	titleForCreate = make(chan string)
	titleForUpdate = make(chan string)
	titleForDelete = make(chan string)
)

func HandleSchedule() {
	for loopTry() != nil {
	}
	println("out loop")
	go copyToBoltdb()
	go copyToHtml()
}

func copyToBoltdb() {
	for {
		ReadMarkdownFiles()
		//键值对比较
		for k, v := range TitleSha256 {
			err := DataBase.View(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte("TitleBucket"))
				vInBoltDb := b.Get([]byte(k))
				if vInBoltDb == nil {
					//没有此 hash ，创建文件
					titleForCreate <- k
				} else if string(vInBoltDb) != v {
					//hash 值冲突，更新文件
					titleForUpdate <- k
				}

				return nil
			})
			if err != nil {
				log.Println(err)
			}

			//遍历
			err = DataBase.View(func(tx *bolt.Tx) error {
				// Assume bucket exists and has keys
				b := tx.Bucket([]byte("TitleBucket"))

				c := b.Cursor()

				for kInBoltDb, vInBoltDb := c.First(); kInBoltDb != nil; kInBoltDb, vInBoltDb = c.Next() {
					if TitleSha256[string(vInBoltDb)] == "" {
						//bolt中有，但是map没有，删除对应文件
						titleForDelete <- string(kInBoltDb)
					}
				}

				return nil
			})
			if err != nil {
				log.Println(err)
			}

		}

		descMaps, err := json.Marshal(DescMaps)
		if err != nil {
			log.Println(err)
		}
		err = DataBase.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("TitleBucket"))
			vInBoltDb := b.Get([]byte("doclist"))
			if vInBoltDb == nil || string(vInBoltDb) != string(descMaps) {
				err := WriteFile(beego.AppConfig.String("outputdir"), "doclist", string(descMaps))
				if err != nil {
					log.Println(err)
				}
				err = DataBase.Update(func(tx *bolt.Tx) error {
					b := tx.Bucket([]byte("TitleBucket"))
					err := b.Put([]byte("doclist"), descMaps)
					return err
				})
			}

			return nil
		})

		time.Sleep(10 * time.Second)
	}

}

func copyToHtml() {
	outdir := beego.AppConfig.String("outputdir")
	for {
		select {
		case msg := <-titleForCreate:
			err := WriteFile(outdir, strings.Replace(msg, ".md", "", 1), MarkdowntoHtml(msg, Files[msg]))
			log.Printf("Write File : %s", strings.Replace(msg, ".md", "", 1))
			if err != nil {
				log.Println(err)
			}
			err = DataBase.Update(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte("TitleBucket"))
				err := b.Put([]byte(msg), []byte(TitleSha256[msg]))
				return err
			})
			if err != nil {
				log.Println(err)
			}
		case msg := <-titleForUpdate:
			err := WriteFile(outdir, strings.Replace(msg, ".md", "", 1), MarkdowntoHtml(msg, Files[msg]))
			log.Printf("Update File : %s", strings.Replace(msg, ".md", "", 1))
			if err != nil {
				log.Println(err)
			}
			err = DataBase.Update(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte("TitleBucket"))
				err := b.Put([]byte(msg), []byte(TitleSha256[msg]))
				return err
			})
			if err != nil {
				log.Println(err)
			}

		case msg := <-titleForDelete:
			err := DeleteFile(outdir, strings.Replace(msg, ".md", "", 1))
			log.Printf("Delete File : %s", strings.Replace(msg, ".md", "", 1))
			if err != nil {
				log.Println(err)
			}
			err = DataBase.Update(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte("TitleBucket"))
				err := b.Delete([]byte(msg))
				if err != nil {
					return err
				}
				return nil
			})
			if err != nil {
				log.Println(err)
			}

		}

	}

}

func loopTry() (err interface{}) {
	defer func() {
		err = recover()
	}()
	time.Sleep(1 * time.Second)
	MarkdowntoHtml("", "")
	return err
}

func MarkdowntoHtml(title string, content string) (strHtml string) {
	var buf bytes.Buffer
	data := make(map[string]string)
	data["Title"] = title
	data["Content"] = content
	err := beego.ExecuteViewPathTemplate(&buf, "page.tpl", beego.BConfig.WebConfig.ViewsPath, data)
	if err != nil {
		return
	}
	strHtml = string(buf.Bytes())
	return
}
