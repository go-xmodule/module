/**
* Created by GoLand
* @file func.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2022/5/11 8:43 上午
* @desc func.go
 */

package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-utils-module/module/global"
	"io/ioutil"
	"os"
	"regexp"
	"runtime/debug"
	"strconv"
	"strings"
)

// Decimal 保留2位小数
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

// MapToSplit map 转换 split
func MapToSplit(mapList map[int]string) (split []string) {
	for _, data := range mapList {
		split = append(split, data)
	}
	return
}

// StringReplace 字符串批量替换
func StringReplace(targetString string, sourceList []string, targetList []string) (string, error) {
	if len(sourceList) != len(targetList) || strings.TrimSpace(targetString) == "" {
		Logger.Error("params error")
		return "", errors.New("params error")
	}
	for key, source := range sourceList {
		targetString = strings.ReplaceAll(targetString, source, targetList[key])
	}
	return targetString, nil
}

// TransInterfaceToStruct 转换interface 到struct
func TransInterfaceToStruct(params interface{}, v interface{}) error {
	jsonData, err := json.Marshal(params)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, v)
	if err != nil {
		return err
	}
	return err
}

// TransInterfaceToMap 转换struct 等结构到map
func TransInterfaceToMap(params interface{}) map[string]interface{} {
	var paramsMap map[string]interface{}
	jsonData, _ := json.Marshal(params)
	_ = json.Unmarshal(jsonData, &paramsMap)
	return paramsMap
}
func ParseFloat64(str string) float64 {
	v, _ := strconv.ParseFloat(str, 64)
	return v
}

// TansToMap 转换json 到map
func TansToMap(paramsStr []byte) map[string]interface{} {
	var params map[string]interface{}
	_ = json.Unmarshal(paramsStr, &params)
	return params
}
func Json(params interface{}) string {
	b, _ := json.Marshal(params)
	return string(b)
}
func JsonString(params interface{}) string {
	b, _ := json.Marshal(params)
	return string(b)
}
func Unmarshal(data string, params interface{}) error {
	return json.Unmarshal([]byte(data), params)
}
func SplitTowString(str string) (string, string, error) {
	temp := strings.Split(str, "@")
	if len(temp) < 2 {
		return "", "", errors.New(global.ParamsError.String())
	}
	return temp[0], temp[1], nil
}

func ReplaceStringByRegex(str, rule, replace string) (string, error) {
	reg, err := regexp.Compile(rule)
	if reg == nil || err != nil {
		return "", errors.New("正则MustCompile错误:" + err.Error())
	}
	return reg.ReplaceAllString(str, replace), nil
}
func Success(status int) bool {
	return global.ErrCode(status) == global.Success
}

func CatchErr(err error, errCode fmt.Stringer, params ...interface{}) bool {
	return HasErr(err, errCode, params...)
}
func HasErr(err error, errCode fmt.Stringer, params ...interface{}) bool {
	if err != nil {
		errMsg := fmt.Sprintf("%s ,err:%s,params:%s,stack:%s", errCode.String(), err.Error(), params, string(debug.Stack()))
		if Logger != nil {
			Logger.Error(errMsg)
		} else {
			fmt.Println(errMsg)
		}
		// go func(err error) {
		// 	_, err = notice.Email.SetContent(errMsg).SendEmail()
		// 	if err != nil {
		// 		Logger.Error("发送错误日志异常", err)
		// 	}
		// }(err)
		return true

	}
	return false
}
func HasWar(err error, errCode fmt.Stringer, params ...interface{}) bool {
	if err != nil {
		errMsg := fmt.Sprintf("%s ,waring:%s,params:%s", errCode.String(), err.Error(), params)
		if Logger != nil {
			Logger.Warning(errMsg)
		} else {
			fmt.Println(errMsg)
		}
		// go func(err error) {
		// 	_, err = notice.Email.SetContent(errMsg).SendEmail()
		// 	if err != nil {
		// 		Logger.Error("发送错误日志异常", err)
		// 	}
		// }(err)
		return true

	}
	return false
}

func CheckErr(err error) bool {
	return err != nil
}
func JsonDisplay(obj interface{}) {
	b, _ := json.Marshal(obj)
	fmt.Println("---------------------------------json obj-------------------------------------")
	var out bytes.Buffer
	_ = json.Indent(&out, b, "", "\t")
	_, _ = out.WriteTo(os.Stdout)
	fmt.Printf("\n")
	fmt.Println("---------------------------------json obj-------------------------------------")
}

func GetFileExtension(fileName string) string {
	temps := strings.Split(fileName, ".")
	return temps[len(temps)-1]
}

func GetImgBase64String(image string) (string, error) {
	bit, err := ioutil.ReadFile(image)
	if err != nil {
		return "", err
	}
	imgStr := base64.StdEncoding.EncodeToString(bit)
	return imgStr, nil
}

func MaxNum(arr []int) (max int, maxIndex int) {
	max = arr[0]
	for i := 0; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
			maxIndex = i
		}
	}
	return max, maxIndex
}

func MinNum(arr []int) (min int, minIndex int) {
	min = arr[0]
	for index, val := range arr {
		if min > val {
			min = val
			minIndex = index
		}
	}
	return min, minIndex
}
