/**
 *@Description
 *@ClassName strategy
 *@Date 2021/5/11 下午8:32
 *@Author ckhero
 */

package strategy_pattern

type Test struct {

}
type IDiscountCalc interface {
	doCalc() int
}

// 策略执行环境
type DiscountPriceCtx struct {
	calc IDiscountCalc
}
const (
	MJ int = iota
	ZK
)
func NewDiscountPriceCtx(discountType int) *DiscountPriceCtx {
	calcMap := map[int]IDiscountCalc{
		MJ: &MJDiscountCalc{},
		ZK: &ZKDiscountCalc{},
	}
	return &DiscountPriceCtx{calc: calcMap[discountType]}
}

func(d DiscountPriceCtx) doCalc() int {
	return d.calc.doCalc()
}
// 满减

type MJDiscountCalc struct {

}

func(MJDiscountCalc) doCalc() int {
	return 100
}

// 折扣

type ZKDiscountCalc struct {

}

func(ZKDiscountCalc) doCalc() int {
	return 200
}

