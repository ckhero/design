/**
 *@Description
 *@ClassName chain
 *@Date 2021/5/11 下午8:55
 *@Author ckhero
 */

package chain_pattern

import "errors"

type IBase interface {
	setNext(next IChain)
	getNext() IChain
}

type IChain interface {
	IBase
	doAuth() (string, error)
}

type Base struct {
	Next IChain
}

func(b Base) setNext(next IChain) {
	b.Next = next
}

func(b Base) getNext() IChain {
	return b.Next
}
//

var authMap = map[int]struct{}{}
// 一级审批
type Level1 struct {
	Base
}

func NewLevel1() IChain {
	return &Level1{}
}

func(l Level1) doAuth() (string, error) {
	_, ok := authMap[1]
	if ok {
		if l.getNext() != nil {
			return l.getNext().doAuth()
		} else {
			return "succ", nil
		}
	}
	return "", errors.New("待一级审批")
}

// 二级审批
type Level2 struct {
	Base
}

func NewLevel2() IChain {
	return &Level2{}
}

func(l Level2) doAuth() (string, error) {
	_, ok := authMap[2]
	if ok {
		if l.getNext() != nil {
			return l.getNext().doAuth()
		} else {
			return "succ", nil
		}
	}
	return "", errors.New("待二级审批")
}
