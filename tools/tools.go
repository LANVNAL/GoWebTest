package tools

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

func ListAll(path string) (fileTarget string, err error) {

	// 获取该路径下的文件列表，并不会像Walk一样遍历子目录
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return "", err
	}
	var fileList []string
	// 过滤掉子目录
	for _, fi := range files {
		if fi.IsDir() {
			// 目录则直接跳过
			continue
		} else {
			fileTarget = filepath.Join(path, fi.Name())
			fileList = append(fileList, fileTarget)
		}
	}
	//转成json
	result, err := json.Marshal(fileList)
	if err != nil{
		return "Convert to JSON Fail", err
	}else{
		return string(result), err
	}
}