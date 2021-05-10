/**
 *@Description
 *@ClassName decorator
 *@Date 2021/5/10 下午9:32
 *@Author ckhero
 */

package decorator_pattern


type IUser interface {
    GetUser() UserInfo
}

type UserInfo struct {
	Name string
	Age int
}

type User struct {
}

func(User) GetUser() UserInfo {
	return UserInfo{
		Name: "name",
		Age:  100,
	}
}

// 装饰器

type Decorator interface {
	IUser
}
// 改名字
type NameDecorator struct {
	user IUser
}

func NewNameDecorator(u IUser) Decorator {
	return &NameDecorator{user: u}
}
func (nd NameDecorator) GetUser() UserInfo {
	userInfo := nd.user.GetUser()
	userInfo.Name = "new name"
	return userInfo
}


// 装饰器 改年龄

type AgeDecorator struct {
	user IUser
}

func NewAgeDecorator(u IUser) Decorator {
	return &AgeDecorator{user: u}
}

func (nd AgeDecorator) GetUser() UserInfo {
	userInfo := nd.user.GetUser()
	userInfo.Age = 200
	return userInfo
}


