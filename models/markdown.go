package models

import "github.com/astaxie/beego"

func ReadMarkdownFiles()  {
	dir := beego.AppConfig.String("inputdir")
	createDirIfNotExist(dir)
	files,err :=  GetAllFile(dir)
	if err != nil {
		panic(err)
	}
	println(files)
}
