package models

import (
	"io/ioutil"
	"os"
	"path"
)

func createDirIfNotExist(dir string)  {
	_, err := os.Stat(dir)
	if err != nil {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

}


func GetAllFile(pathname string,predicate func (string) bool) (files map[string][]byte ,err error) {
	if files ==nil {
		files = make(map[string][]byte)
	}
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if ! fi.IsDir() && predicate(fi.Name()) {
			file, err := ioutil.ReadFile(path.Join(pathname,fi.Name()))
			if err != nil {
				panic(err)
			}
			files[fi.Name()] = []byte(file)
		}else {
			files,err =  GetAllFile(pathname,predicate)
			if err != nil {
				panic(err)
			}
		}
	}
	return
}