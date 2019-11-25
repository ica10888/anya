package models

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/boltdb/bolt"
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
			DataBase.View(func(tx *bolt.Tx) error {
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

			//遍历
			DataBase.View(func(tx *bolt.Tx) error {
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

		}

		time.Sleep(10 * time.Second)
	}

}

func copyToHtml() {
	for {
		select {
		case msg := <-titleForCreate:
			fmt.Println(msg)

		case msg := <-titleForUpdate:
			fmt.Println(msg)

		case msg := <-titleForDelete:
			fmt.Println(msg)

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
	beego.ExecuteViewPathTemplate(&buf, "page.tpl", beego.BConfig.WebConfig.ViewsPath, data)
	strHtml = string(buf.Bytes())
	return
}
