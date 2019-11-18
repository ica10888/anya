package models

import (
	"bytes"
	"github.com/astaxie/beego"
	"github.com/shurcooL/github_flavored_markdown"
	"gopkg.in/yaml.v2"
	"regexp"
	"strings"
)

var Markdowns map[string]DocPage

type DocDescription struct {
	Layout   string   `yaml:"layout"  json:"layout"`
	Title    string   `yaml:"title"  json:"title"`
	Subtitle string   `yaml:"subtitle"  json:"subtitle"`
	Date     string   `yaml:"date"  json:"date"`
	Author   string   `yaml:"author"  json:"author"`
	Tags     []string `yaml:"tags"  json:"tags"`
}

type DocPage struct {
	DocDesc    DocDescription
	DocContent string
}

var Content string

func ReadMarkdownFiles() {
	dir := beego.AppConfig.String("inputdir")
	createDirIfNotExist(dir)
	files, err := GetAllFile(dir, func(s string) bool {
		regex, err := regexp.Match(`.*.md`, []byte(s))
		if err != nil {
			panic(err)
		}
		return regex
	})
	if err != nil {
		panic(err)
	}

	for k, v := range files {
		sha256 := GetSHA256HashCode(v)
		println(k)
		println(sha256)
	}
	MarkdowntoHtml()

	Markdowns = make(map[string]DocPage)
	for k, v := range files {
		docMarkdown := strings.SplitN(v, "---", 3)
		var dd DocDescription
		err := yaml.Unmarshal([]byte(docMarkdown[1]), &dd)
		if err != nil {
			panic(err)
		}
		Markdowns[strings.Replace(k, ".md", "", 1)] = DocPage{dd,
			string(toMarkdown([]byte(docMarkdown[2])))}

	}

}

func toMarkdown(content []byte) (output []byte) {
	output = github_flavored_markdown.Markdown(content)
	return
}

func MarkdowntoHtml() {
	var buf bytes.Buffer
	beego.ExecuteViewPathTemplate(&buf, "page.tpl", "", "")

}
