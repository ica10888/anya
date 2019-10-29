package models

import (
	"github.com/astaxie/beego"
	"github.com/shurcooL/github_flavored_markdown"
	"regexp"
)

var Content string

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

	markdowns := make(map[string][]byte)
	for k, v := range files {
		markdowns[k] = toMarkdown(v)
		s:= string(toMarkdown(v))
		Content = s
	}
	println(markdowns)
}


func toMarkdown (content []byte) (output []byte)  {
	output = github_flavored_markdown.Markdown(content)
	return
}


