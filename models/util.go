package models

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"os"
	"path"
)

func createDirIfNotExist(dir string) {
	_, err := os.Stat(dir)
	if err != nil {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

}

func GetAllFile(pathname string, predicate func(string) bool) (files map[string]string, err error) {
	if files == nil {
		files = make(map[string]string)
	}
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if !fi.IsDir() && predicate(fi.Name()) {
			file, err := ioutil.ReadFile(path.Join(pathname, fi.Name()))
			if err != nil {
				panic(err)
			}
			files[fi.Name()] = string(file)
		} else {
			files, err = GetAllFile(pathname, predicate)
			if err != nil {
				panic(err)
			}
		}
	}
	return
}

func WriteFile(pathname string, fileName string, content string) (err error) {
	paths := path.Join(pathname, fileName)
	f, err := os.Create(paths)
	defer f.Close()
	if err != nil {
		return
	} else {
		_, err = f.Write([]byte(content))
	}
	return
}

func DeleteFile(pathname string, fileName string) (err error) {
	paths := path.Join(pathname, fileName)
	err = os.Remove(paths)
	return
}

func GetSHA256HashCode(message string) (hashCode string) {
	hash := sha256.New()
	hash.Write([]byte(message))
	bytes := hash.Sum(nil)
	hashCode = hex.EncodeToString(bytes)
	return
}
