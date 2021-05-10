/**
 *@Description
 *@ClassName status
 *@Date 2021/5/10 上午10:10
 *@Author ckhero
 */

package status_pattern

import (
	"errors"
)

const (
	INIT int = iota
	COMMIT
	SUCC
	CANCEL
)


type StatusHanlde interface {
	Commit() (bool, error)
	Succ() (bool, error)
	Cancel() (bool, error)
}

// 初始化状态
type InitStatus struct {
	
}
func newInitStatus() StatusHanlde {
	return &InitStatus{}
}
func (s *InitStatus) Commit() (bool, error) {
	return true, nil
}

func (s *InitStatus) Succ() (bool, error) {
	return false, errors.New("init to succ fail")
}

func (s *InitStatus) Cancel() (bool, error) {
	return false, errors.New("init to cancel fail")
}

type CommitStatus struct {

}
func newCommitStatus() StatusHanlde {
	return &CommitStatus{}
}
func (s *CommitStatus) Commit() (bool, error) {
	return false, errors.New("already commit")
}

func (s *CommitStatus) Succ() (bool, error) {
	return true, nil
}

func (s *CommitStatus) Cancel() (bool, error) {
	return true, nil
}


// 成功状态
type SuccStatus struct {

}
func newCuccStatus() StatusHanlde {
	return &SuccStatus{}
}
func (s *SuccStatus) Commit() (bool, error) {
	return false, errors.New("succ to commit fail")
}

func (s *SuccStatus) Succ() (bool, error) {
	return false, errors.New("succ aleady")
}

func (s *SuccStatus) Cancel() (bool, error) {
	return false, errors.New("succ to cancel fail")
}

var statusHandleMap = map[int]func() StatusHanlde {
	INIT : newInitStatus,
	COMMIT: newCommitStatus,
	SUCC: newCuccStatus,
}
func NewStatusHandle(currStatus int) StatusHanlde {
	return statusHandleMap[currStatus]()
}