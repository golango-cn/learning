package model

import (
	"fmt"
	. "github.com/ahmetb/go-linq"
	"strings"
	"time"
)

type Good struct {
	Id       int32   `json:"id"`
	Name     string  `json:"name"`
	Weight   float64 `json:"weight"`
	Price    float64 `json:"price"`
	CreateAt int64   `json:"create_at"`
}

func (g *Good) Gets(keys string, pageIndex, pageSize int) ([]Good, int32, error) {

	var goods []Good
	total := 1000

	for i := 1000; i <= total+1000; i++ {
		goods = append(goods, Good{
			Id:       int32(i),
			Name:     fmt.Sprintf("商品%d", i),
			Weight:   float64(i),
			Price:    float64(i),
			CreateAt: time.Now().Unix(),
		})
	}

	From(goods).Where(func(i interface{}) bool {
		if len(keys) > 0 {
			return strings.Contains(i.(Good).Name, keys)
		}
		return true
	}).Skip((pageIndex - 1) * pageSize).Take(pageSize).ToSlice(&goods)
	return goods, int32(total), nil
}
