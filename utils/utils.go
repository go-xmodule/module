/**
 * Created by goland.
 * @file   utils.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2023/2/7 12:33
 * @desc   utils.go
 */

package utils

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
)

func CompareHashAndPassword(e string, p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(e), []byte(p))
	if err != nil {
		return false, err
	}
	return true, nil
}

// FilterSearchField 过滤掉Pagination中的 pageIndex 和 pageSize 字段
func FilterSearchField(s interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	jsonStr, _ := json.Marshal(s)
	_ = json.Unmarshal(jsonStr, &m)
	delete(m, "PageIndex")
	delete(m, "PageSize")
	for key, value := range m {
		if value == nil || value == "" || value == 0 {
			delete(m, key)
		}
	}
	return m
}
