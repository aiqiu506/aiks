package utils

import (
	"errors"
	"reflect"
)

func MapToStruct(mmap map[interface{}]interface{},structure interface{}) (err error){
	defer func() {
		if errs,ok := recover().(error); ok {
			err = errs
		}
	}()
	ptp := reflect.TypeOf(structure)
	pv := reflect.ValueOf(structure)
	switch ptp.Kind() {
	case reflect.Ptr:
		if ptp.Elem().Kind() == reflect.Struct {
			break
		}else{
			return errors.New("需要*struct类型，却传入*"+ptp.Elem().Kind().String()+"类型")
		}
	default:
		return errors.New("需要*struct类型，却传入"+ptp.Kind().String()+"类型")
	}
	tp := ptp.Elem()
	v := pv.Elem()
	num := tp.NumField()
	for i := 0 ; i < num ; i++ {
		name := tp.Field(i).Name
		//fmt.Println(name)
		tag := tp.Field(i).Tag.Get("map")
		if len(tag) != 0 {
			name = tag
		}
		value,ok := mmap[name]
		if !ok {
			continue
		}
		//能够设置值，且类型相同
		if v.Field(i).CanSet(){
			if v.Field(i).Type() == reflect.TypeOf(value){
				v.Field(i).Set(reflect.ValueOf(value))
			}else{
				continue
			}
		}else {
			continue
		}
	}
	return nil
}