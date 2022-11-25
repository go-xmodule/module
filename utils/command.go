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
	"context"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type Output func(result string)

func ExecuteWithContext(ctx context.Context, cmd string, output Output) error {
	c := exec.CommandContext(ctx, "bash", "-c", cmd) // mac linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	reader := bufio.NewReader(stdout)
	err = c.Start()
	for {
		// 其实这段去掉程序也会正常运行，只是我们就不知道到底什么时候Command被停止了，而且如果我们需要实时给web端展示输出的话，这里可以作为依据 取消展示
		select {
		// 检测到ctx.Done()之后停止读取
		case <-ctx.Done():
			if ctx.Err() != nil {
				output("程序出现错误: " + ctx.Err().Error())
			} else {
				output("程序被终止")
			}
			return nil
		default:
			outputStr, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return nil
			}
			res := strings.TrimSpace(strings.Replace(outputStr, "\\n", "", -1))
			if len(res) > 0 {
				output(res)
			}
		}
	}
}

// ExecuteCommand 执行命令
func ExecuteCommand(executeCmd string, result Output) error {
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

func getOutput(reader *bufio.Reader, callback Output) string {
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
