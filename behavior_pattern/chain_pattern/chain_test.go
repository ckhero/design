/**
 *@Description
 *@ClassName chain_test
 *@Date 2021/5/11 下午9:09
 *@Author ckhero
 */

package chain_pattern

import (
	"fmt"
	"testing"
)

func TestNewLevel1(t *testing.T) {
	chain := NewLevel1()
	chain.setNext(NewLevel2())
	fmt.Println(chain.doAuth())
	authMap[1] = struct{}{}
	fmt.Println(chain.doAuth())
	authMap[1] = struct{}{}
	fmt.Println(chain.doAuth())
}
