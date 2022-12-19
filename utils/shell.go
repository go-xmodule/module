/**
 * Created by goland.
 * @file   shell.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/12/14 15:32
 * @desc   shell.go
 */

package utils

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// GetCmdPath 获取命令路径
func GetCmdPath(cmd string) (string, error) {
	c := exec.Command("bash", "-c", "whereis "+cmd)
	output, err := c.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fields := strings.Fields(string(output))
	if len(fields) > 1 {
		return fields[1], nil
	}
	return "", errors.New("not command found")
}
