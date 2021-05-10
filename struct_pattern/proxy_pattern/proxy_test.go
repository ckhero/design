/**
 *@Description
 *@ClassName proxy_test
 *@Date 2021/5/10 下午9:16
 *@Author ckhero
 */

package proxy_pattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCacheProxy_GetUser(t *testing.T) {

	u := &User{}
	assert.Equal(t, "user info", u.GetUser())

	log := NewLogProxy(u)
	assert.Equal(t, "[log] user info", log.GetUser())

	cache := NewCahceProxy(log)
	assert.Equal(t, "[cache] [log] user info", cache.GetUser())

}