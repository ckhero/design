/**
 *@Description
 *@ClassName status_test
 *@Date 2021/5/10 上午10:29
 *@Author ckhero
 */

package status_pattern

import (
	"fmt"
	"testing"
)

func TestNewStatusHandle(t *testing.T) {
	sh := NewStatusHandle(INIT)
	fmt.Println(sh.Commit())
	fmt.Println(sh.Succ())

	sh = NewStatusHandle(SUCC)
	fmt.Println(sh.Commit())
	fmt.Println(sh.Succ())
}
