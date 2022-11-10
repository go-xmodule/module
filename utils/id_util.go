package utils

import (
	"strings"

	"github.com/google/uuid"
)

// GetUUIdByTime 根据时间生成UUID true去除“-”，false不去除
func GetUUIdByTime(flag bool) (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	if flag {
		return strings.Replace(id.String(), "-", "", -1), nil
	}
	return id.String(), err
}

// IdUUIdByRand V4 基于随机数 true去除“-”，false不去除
func IdUUIdByRand(flag bool) string {
	u4 := uuid.New()
	if flag {
		return strings.Replace(u4.String(), "-", "", -1)
	}
	return u4.String()
}
