package leetcode

import (
	"io/ioutil"
	"os"
	"strings"
)

//RenameFiles 文件重命名
func RenameFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			RenameFiles(dirPth + PthSep + fi.Name())
		} else {
			filePath := dirPth + PthSep
			fileName := fi.Name()
			ok := strings.HasSuffix(fileName, "go")
			if ok {
				if strings.HasSuffix(fileName, "_test.go") {
					os.Rename(filePath+fileName, filePath+"code_test.go")
				} else {
					os.Rename(filePath+fileName, filePath+"code.go")
				}
			}
		}
	}

	return files, nil
}
