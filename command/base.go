/**
 * Created by PhpStorm.
 * @file   base.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/20 00:35
 * @desc   base.go
 */

package command

import "github.com/urfave/cli/v2"

// Inter 命令行接口
type Inter interface {
	Command() *cli.Command
	Action(ctx *cli.Context) error
}

// BaseCommand 基础命令
type BaseCommand struct {
}
