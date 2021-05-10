/**
 *@Description
 *@ClassName proxy
 *@Date 2021/5/10 下午9:13
 *@Author ckhero
 */

package proxy_pattern

import "fmt"

type IUser interface {
	GetUser() string
}

type User struct {

}

func (User) GetUser() string {
	return "user info"
}
// 日志代理
type LogProxy struct {
	user IUser
}

func NewLogProxy(user IUser) IUser {
	return &LogProxy{user: user}
}

func (l LogProxy) GetUser() string{
	return fmt.Sprintf("[%s] %s", "log", l.user.GetUser())
}
// 缓存代理
func NewCahceProxy(proxy IUser) IUser {
	return &CacheProxy{logProxy: proxy}
}

type CacheProxy struct {
	logProxy IUser
}

func (l CacheProxy) GetUser() string {
	return fmt.Sprintf("[%s] %s", "cache", l.logProxy.GetUser())
}


