/**
 * Created by PhpStorm.
 * @file   math.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/2 16:56
 * @desc   math.go
 */

package utils

// Average 球平局数
func Average(xs []float64) (avg float64) {
	sum := 0.0
	if len(xs) == 0 {
		avg = 0
	} else {
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs))
	}
	return
}
