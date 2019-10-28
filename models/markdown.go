package models

import (
	"github.com/astaxie/beego"
	"regexp"
)

func ReadMarkdownFiles()  {
	dir := beego.AppConfig.String("inputdir")
	createDirIfNotExist(dir)
	files,err :=  GetAllFile(dir, func(s string) bool {
		regex,err := regexp.Match(`.*.md`, []byte(s))
		if err != nil {
			panic(err)
		}
		return regex
	})
	if err != nil {
		panic(err)
	}
	println(files)
}
