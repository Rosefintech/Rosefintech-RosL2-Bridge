/*
@Author :
@Group :Rosefintech
@Description :
@Data: 2020/6/3 15:43
*/

package tools

import (
	"io/ioutil"
	"os"
)

func IsFile(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return !fi.IsDir(), nil
}

func IsDir(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fi.IsDir(), nil
}


func DeleteFile(path string) (err error) {
	err = os.Remove(path)
	return
}

func GetLocalPwd() (string, error) {
	dir, err := os.Getwd()
	return dir, err
}
func Open(path string) (file *os.File, err error) {
	file, err = os.Open(path)
	return
}

func OpenFile(path string) (file *os.File, err error) {
	file, err = os.OpenFile(path, os.O_RDWR, 0600)
	return
}

func CreatFile(path string) (file *os.File, err error) {
	file, err = os.Create(path)
	return
}

func CreatDir(name string) (err error) {
	err = os.Mkdir(name, os.ModePerm)
	return
}

func DeleteAll(path string) (err error) {
	err = os.RemoveAll(path)
	return
}

func ReadFile(path string) (data []byte, err error) {
	data, err = ioutil.ReadFile(path)
	return
}
