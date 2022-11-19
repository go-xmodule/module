package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dlclark/regexp2"
	"reflect"
	"strings"
)

type EncryptUtil struct {
}

// Md5 MD5算法.
func Md5(input interface{}) (string, error) {
	params := ""
	inputType := reflect.TypeOf(input).String()
	if inputType == "string" {
		params = input.(string)
	} else {
		in, err := json.Marshal(input)
		if err != nil {
			return "", err
		}
		params = string(in)
	}
	hash := md5.New()
	_, err := hash.Write([]byte(params))
	if err != nil {
		return "", err
	}
	result := hash.Sum(nil)
	return fmt.Sprintf("%x", result), nil
}

// SHA1 sha1加密 encryption
func SHA1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return strings.ToUpper(hex.EncodeToString(o.Sum(nil)))
}

func encode() {
	data := "hello world12345!?$*&()'-@~"
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc) // aGVsbG8gd29ybGQxMjM0NSE/JComKCknLUB+
}

func MatchPassword(str string) error {
	if str == "" {
		return errors.New("密码不能为空")
	}
	expr := `^(?![0-9a-zA-Z]+$)(?![a-zA-Z!@#$%^&*]+$)(?![0-9!@#$%^&*]+$)[0-9A-Za-z!@#$%^&*]{8,16}$`
	reg, _ := regexp2.Compile(expr, 0)
	m, _ := reg.FindStringMatch(str)
	if m != nil {
		return nil
	}
	return errors.New("密码包含至少一位数字，字母和特殊字符,且长度8-16")
}
