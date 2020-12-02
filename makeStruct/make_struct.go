package makeStruct

import (
	"reflect"
)

func MakeStruct(target interface{}) {
	if target == nil {
		return
	}
	rType := reflect.TypeOf(target)
	rVal := reflect.ValueOf(target)

	if rType.Kind() == reflect.Ptr {
		rType = reflect.TypeOf(target).Elem()
		rVal = reflect.ValueOf(target).Elem()
	}

	for i := 0; i < rType.NumField(); i++ {
		if !rVal.Field(i).CanAddr() {
			continue
		}
		t := rType.Field(i)
		f := rVal.Field(i)
		switch t.Type.Kind() {
		case reflect.Slice:
			l := f.Len()
			if l > 0 {
				for index := 0; index < l; index++ {
					MakeStruct(f.Index(index).Interface())
				}
			} else {
				f.Set(reflect.MakeSlice(reflect.SliceOf(t.Type.Elem()), 0, 0))
			}
		case reflect.Struct:
			//非指针结构体，取出它的地址递归
			MakeStruct(rVal.Field(i).Addr().Interface())
		case reflect.Ptr:
			//如果为空指针，就new一个新的该类型赋值给它
			if rVal.Field(i).IsNil() {
				rVal.Field(i).Set(reflect.New(rType.Field(i).Type.Elem()))
			}
			//如果真实类型非结构体，那么无需进一步处理
			if rType.Field(i).Type.Elem().Kind() != reflect.Struct {
				continue
			}
			MakeStruct(rVal.Field(i).Interface())
		}
	}
}
