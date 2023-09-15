/**
 * Created by goland.
 * @file   utils.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2023/2/7 12:33
 * @desc   utils.go
 */

package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/x-module/module/config"
	"github.com/x-module/utils/global"
	"github.com/x-module/utils/utils/xlog"
	"go-micro.dev/v4/metadata"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// GetImageUrl 格式化&拼接图片地址
func GetImageUrl(config config.ApiServer, path string) string {
	return fmt.Sprintf("%s://%s:%d/%s", config.Protocol, config.Domain, config.Port, path)
}

func Logger(ctx context.Context) *logrus.Entry {
	res, _ := metadata.FromContext(ctx)
	return xlog.Logger.WithFields(logrus.Fields{
		"playerId":  res["Playerid"],
		"requestId": res["Requestid"],
		"server":    "player",
	})
}

// HasQueryErr 数据库查询异常
func HasQueryErr(ctx context.Context, err error, errCode fmt.Stringer) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		msg := fmt.Sprintf("%s desc:%s", errCode.String(), global.NoRecordErr.String())
		Logger(ctx).WithField(global.ErrField, err).Warn(msg)
	} else {
		Logger(ctx).WithField(global.ErrField, err).Error(errCode.String())
	}
	return true
}
func HasErr(ctx context.Context, err error, errCode fmt.Stringer) bool {
	if err != nil {
		Logger(ctx).WithField("err", err).Error(errCode.String())
		return true
	}
	return false
}

// HasWar 通用异常处理
func HasWar(ctx context.Context, err error, errCode fmt.Stringer) bool {
	if err != nil {
		Logger(ctx).WithField("err", err).Warn(errCode.String())
		return true
	}
	return false
}

func CatchErr(ctx context.Context, err error, errCode fmt.Stringer) bool {
	return HasErr(ctx, err, errCode)
}

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
