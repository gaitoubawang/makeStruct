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
			// got the addr of the struct to recursive
			MakeStruct(rVal.Field(i).Addr().Interface())
		case reflect.Ptr:
			// if it's nil, new value for it
			if rVal.Field(i).IsNil() {
				rVal.Field(i).Set(reflect.New(rType.Field(i).Type.Elem()))
			}
			// if ptr points the object type is not struct , continue
			if rType.Field(i).Type.Elem().Kind() != reflect.Struct {
				continue
			}
			MakeStruct(rVal.Field(i).Interface())
		}
	}
}
