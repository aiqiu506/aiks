package utils

import (
	"errors"
	"reflect"
)

/**
 判断某元素是否存在数姐中
@params matherArray 被查找数组
@params sonEle  要判断的元素
*/

func InArray(matherArray interface{}, sonEle interface{}) (int, error) {
	mt := reflect.TypeOf(matherArray)
	mv := reflect.ValueOf(matherArray)
	v := reflect.ValueOf(sonEle)
	if mt.Kind() != reflect.Slice && mt.Kind() != reflect.Array {
		return -1, errors.New("matherArray is not Slice type")
	}
	for i, lens := 0, mv.Len(); i < lens; i++ {
		if reflect.DeepEqual(mv.Index(i).Interface(), v.Interface()) {
			return i, nil
		}
	}
	return -1, nil
}