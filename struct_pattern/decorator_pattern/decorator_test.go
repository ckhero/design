/**
 *@Description
 *@ClassName decorator_test
 *@Date 2021/5/10 下午9:43
 *@Author ckhero
 */

package decorator_pattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_GetUser(t *testing.T) {
	u := User{}
	assert.Equal(t, UserInfo{
		Name: "name",
		Age:  100,
	}, u.GetUser())

	age := NewAgeDecorator(u)
	assert.Equal(t, UserInfo{
		Name: "name",
		Age:  200,
	}, age.GetUser())

	name := NewNameDecorator(u)
	assert.Equal(t, UserInfo{
		Name: "new name",
		Age:  100,
	}, name.GetUser())
}
