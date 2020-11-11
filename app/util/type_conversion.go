package util

import "reflect"

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj) // 获取 obj 的类型信息
	v := reflect.ValueOf(obj)
	if t.Kind() == reflect.Ptr { // 如果是指针，则获取其所指向的元素
		t = t.Elem()
		v = v.Elem()
	}
	var data = make(map[string]interface{})
	if t.Kind() == reflect.Struct { // 只有结构体可以获取其字段信息
		for i := 0; i < t.NumField(); i++ {
			data[t.Field(i).Name] = v.Field(i).Interface()
		}
	}
	return data
}
