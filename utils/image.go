/**
* Created by GoLand
* @file image.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2022/2/20 11:30 上午
* @desc image.go
 */

package utils

import (
	"bytes"
	"github.com/disintegration/imaging"
	"io/ioutil"
)

type ImageUtil struct {
	width     int
	height    int
	sourceImg string
	targetImg string
}

func (i *ImageUtil) SetImage(sourceImg string, targetImg string) *ImageUtil {
	i.sourceImg = sourceImg
	i.targetImg = targetImg
	return i
}
func (i *ImageUtil) SetSize(width int, height int) *ImageUtil {
	i.height = height
	i.width = width
	return i
}
func (i *ImageUtil) Resize() error {
	imgData, err := ioutil.ReadFile(i.sourceImg)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(imgData)
	image, err := imaging.Decode(buf)
	if err != nil {
		return err
	}
	image = imaging.Resize(image, i.width, i.height, imaging.Lanczos)
	err = imaging.Save(image, i.targetImg)
	return err
}
