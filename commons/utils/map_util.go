package utils

import "reflect"

// Keys 获取键列表
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// Values 获取值列表
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

// Equal 相等
func Equal[M1, M2 ~map[K]V, K, V comparable](m1 M1, m2 M2) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}

// EqualFunc 相等
func EqualFunc[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 any](m1 M1, m2 M2, eq func(V1, V2) bool) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || !eq(v1, v2) {
			return false
		}
	}
	return true
}

// Clear 清空
func Clear[M ~map[K]V, K comparable, V any](m M) {
	for k := range m {
		delete(m, k)
	}
}

// Clone 克隆
func Clone[M ~map[K]V, K comparable, V any](m M) M {
	if m == nil {
		return nil
	}
	r := make(M, len(m))
	for k, v := range m {
		r[k] = v
	}
	return r
}

// Concat 连接
func Concat[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2) {
	for k, v := range src {
		dst[k] = v
	}
}

// DeleteFunc 删除
func DeleteFunc[M ~map[K]V, K comparable, V any](m M, del func(K, V) bool) {
	for k, v := range m {
		if del(k, v) {
			delete(m, k)
		}
	}
}

// StructToMap 结构体转换map
func StructToMap(obj any) map[string]interface{} {
	data := make(map[string]interface{})
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	for i := 0; i < objT.NumField(); i++ {
		fileName, ok := objT.Field(i).Tag.Lookup("json")
		if ok {
			data[fileName] = objV.Field(i).Interface()
		} else {
			data[objT.Field(i).Name] = objV.Field(i).Interface()
		}
	}
	return data
}

// StructToSrcMap 结构体转换来源map
func StructToSrcMap(obj any, srcMap map[string]interface{}) map[string]interface{} {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	for i := 0; i < objT.NumField(); i++ {
		fileName, ok := objT.Field(i).Tag.Lookup("json")
		if ok {
			srcMap[fileName] = objV.Field(i).Interface()
		} else {
			srcMap[objT.Field(i).Name] = objV.Field(i).Interface()
		}
	}
	return srcMap
}
