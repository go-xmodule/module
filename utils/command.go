/**
* Created by GoLand
* @file main.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2022/07/19 6:19 下午
* @desc   执行系统命令
 */

package utils

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type Response func(result string)

// ExecuteCommand 执行命令
func ExecuteCommand(executeCmd string, result Response) error {
	commands := strings.Fields(executeCmd)
	cmd := exec.Command(commands[0], commands[1:]...)
	cmd.Stdin = os.Stdin
	var wg sync.WaitGroup
	wg.Add(2)
	// 捕获标准输出
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	readout := bufio.NewReader(stdout)
	go func() {
		defer wg.Done()
		getOutput(readout, result)
	}()
	// 捕获标准错误
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	readErr := bufio.NewReader(stderr)
	go func() {
		defer wg.Done()
		getOutput(readErr, result)
	}()
	// 执行命令
	err = cmd.Run()
	if err != nil {
		return err
	}
	wg.Wait()
	return nil
}

func getOutput(reader *bufio.Reader, callback Response) string {
	var sumOutput string // 统计屏幕的全部输出内容
	outputBytes := make([]byte, 200)
	for {
		n, err := reader.Read(outputBytes) // 获取屏幕的实时输出(并不是按照回车分割，所以要结合sumOutput)
		if err != nil {
			if err == io.EOF {
				break
			}
			sumOutput += err.Error()
		}
		output := string(outputBytes[:n])
		callback(output)
		sumOutput += output
	}
	callback("--end--")
	return sumOutput
}
