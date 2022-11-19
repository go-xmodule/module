/**
 * Created by PhpStorm.
 * @file   upload.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/20 02:10
 * @desc   upload.go
 */

package config

// Upload 图片设置
type Upload struct {
	UploadDir string `yaml:"uploadDir"`
	TempDir   string `yaml:"tempDir"`
}
