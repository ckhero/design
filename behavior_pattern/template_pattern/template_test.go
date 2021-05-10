/**
 *@Description
 *@ClassName template_test
 *@Date 2021/5/10 上午11:10
 *@Author ckhero
 */

package template_pattern

import (
	"fmt"
	"testing"
)

func TestNewTemplate(t *testing.T) {
	base := BaseTemplate{}
	fmt.Println(base.Exec(NewJD()))
	fmt.Println(base.Exec(NewTaobao()))
}