/**
 *@Description
 *@ClassName strategy_test
 *@Date 2021/5/11 下午8:41
 *@Author ckhero
 */

package strategy_pattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDiscountPriceCtx(t *testing.T) {
	mj := NewDiscountPriceCtx(MJ)
	assert.Equal(t, 100, mj.doCalc())

	zk := NewDiscountPriceCtx(ZK)
	assert.Equal(t, 200, zk.doCalc())
}
