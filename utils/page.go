/**
* Created by GoLand
* @file utils.go
* @version: 1.0.0
* @author 李锦 <Lijin@cavemanstudio.net>
* @date 2022/2/2 9:50 下午
* @desc utils.go
 */

package utils

import (
	"math"
)

type PageUtil struct{}

type PageInfo struct {
	PrevPage   int
	CurrPage   int
	NextPage   int
	Pages      []int
	TotalCount int
	TotalPage  int
	Offset     int
	EndOffset  int
}
type PageData struct {
	PageInfo PageInfo
	DataList interface{}
}

func (u *PageUtil) Paginator(currentPage, prePage int, totalCount int64) PageInfo {
	var prevPage int // 前一页地址
	var nextPage int // 后一页地址
	// 根据totalCount总数，和prePage每页数量 生成分页总数
	totalPage := int(math.Ceil(float64(totalCount) / float64(prePage))) // page总数
	if currentPage > totalPage {
		currentPage = totalPage
	}
	if currentPage <= 0 {
		currentPage = 1
	}
	var pages []int
	switch {
	case currentPage >= totalPage-5 && totalPage > 5: // 最后5页
		start := totalPage - 5 + 1
		prevPage = currentPage - 1
		nextPage = int(math.Min(float64(totalPage), float64(currentPage+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case currentPage >= 3 && totalPage > 5:
		start := currentPage - 3 + 1
		pages = make([]int, 5)
		prevPage = currentPage - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		prevPage = currentPage - 1
		nextPage = currentPage + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalPage))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		prevPage = int(math.Max(float64(1), float64(currentPage-1)))
		nextPage = currentPage + 1
	}
	return PageInfo{
		PrevPage:   prevPage,
		NextPage:   nextPage,
		TotalPage:  totalPage,
		CurrPage:   currentPage,
		Pages:      pages,
		TotalCount: int(totalCount),
	}
}
