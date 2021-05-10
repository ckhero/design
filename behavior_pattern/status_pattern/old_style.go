/**
 *@Description
 *@ClassName old_style
 *@Date 2021/5/10 上午10:19
 *@Author ckhero
 */

package status_pattern

import "errors"

func OldUpdateStatus(oldStatus, currStatus int) (bool, error) {
	if oldStatus == INIT {
		if currStatus == COMMIT {
			return true, errors.New("init to succ fail")
		}
		if currStatus == SUCC {
			return false, errors.New("init to succ fail")
		}
		if currStatus == CANCEL {
			return false, errors.New("init to cancel fail")
		}
	}

	if oldStatus == COMMIT {
		if currStatus == INIT {
			return false, errors.New("commit to init fail")
		}
		if currStatus == SUCC {
			return true, nil
		}
		if currStatus == CANCEL {
			return true, nil
		}
	}
	return false, errors.New("invalid status")
}
