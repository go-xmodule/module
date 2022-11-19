/**
 * Created by PhpStorm.
 * @file   dir.go
 * @author 李锦 <Ljin@cavemanstudio.net>
 * @date   2022/7/22 17:12
 * @desc   dir.go
 */

package utils

import (
	"fmt"
	"io/ioutil"
)

// ListDirNames return all file names in the path
func ListDirNames(path string) ([]string, error) {
	fs, err := ioutil.ReadDir(path)
	if err != nil {
		return []string{}, err
	}
	sz := len(fs)
	if sz == 0 {
		return []string{}, nil
	}
	var res []string
	for i := 0; i < sz; i++ {
		if fs[i].IsDir() {
			res = append(res, fs[i].Name())
		}
	}
	return res, nil
}

// FormatFileSize 字节的单位转换 保留两位小数
func FormatFileSize(fileSize float64) (size string) {
	if fileSize < 1024 {
		// return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", fileSize/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", fileSize/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", fileSize/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", fileSize/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", fileSize/float64(1024*1024*1024*1024))
	} else { // if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", fileSize/float64(1024*1024*1024*1024*1024))
	}
}
