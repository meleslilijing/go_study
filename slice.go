package week1

import "errors"

var ErrIndexOutOfRange = errors.New("下标超出范围")

// DeleteAt 删除指定位置的元素
// 如果下标不是合法的下标，返回 ErrIndexOutOfRange
func DeleteAt[Type any](s []Type, idx int) ([]Type, error) {
	if idx >= len(s) || idx < 0 {
		return nil, ErrIndexOutOfRange
	}

	s1 := make([]Type, len(s), cap(s)-1)

	for i := 0; i < idx; i++ {
		s1[i] = s[i]
	}

	for i := idx; i < len(s); i++ {
		s1[i-1] = s[i]
	}

	return s1, nil
}
