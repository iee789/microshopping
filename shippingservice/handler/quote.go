package handler

import (
	"fmt"
	"math"
)

// 货币描述
type Quote struct {
	Dollars uint32
	Cents   uint32
}

// 报价描述
func (q Quote) String() string {
	return fmt.Sprintf("$%d.%d", q.Dollars, q.Cents)
}

// 根据商品数量创建报价
func CreateQuoteFromCount(count int) Quote {
	return CreateQuoteFromFloat(8.99)
}

// 创建报价
func CreateQuoteFromFloat(value float64) Quote {
	units, fraction := math.Modf(value) // 把value分解成整数和小数部分
	return Quote{
		uint32(units),
		// 将fraction乘以100，这样小数点后两位的值就会被放大到整数范围内
		// 然后使用 math.Trunc函数截断这个放大后的浮点数，得到一个整数：代表了原始浮点数小数点后两位的值
		uint32(math.Trunc(fraction * 100)),
	}
}
