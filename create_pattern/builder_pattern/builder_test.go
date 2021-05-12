/**
 *@Description
 *@ClassName builder_test
 *@Date 2021/5/12 下午10:08
 *@Author ckhero
 */

package builder_pattern

import (
	"fmt"
	"testing"
)

func TestNewBuilder(t *testing.T) {
	d := &Director{build: NewBuilder()}
	fmt.Println(d.Create01().GetDetail())
	fmt.Println(d.Create02().GetDetail())
}
