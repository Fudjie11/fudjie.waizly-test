package util

import (
	"errors"
	"reflect"
)

var AllIntReflectKind = []reflect.Kind{reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64}
var AllFloatReflectKind = []reflect.Kind{reflect.Float32, reflect.Float64}

// ref: https://vitaneri.com/posts/check-for-nil-interface-in-go
func IsNilableReflectKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
		return true
	default:
		return false
	}
}

// ref: https://vitaneri.com/posts/check-for-nil-interface-in-go
func IsNilReflectValue(val reflect.Value) bool {
	if !val.IsValid() {
		return true
	}

	return IsNilableReflectKind(val.Kind()) && val.IsNil()
}

func GetUnderlyingReflectValueInfo(oriRef reflect.Value) (ref reflect.Value, refType reflect.Type, refKind reflect.Kind, refValue interface{}) {
	ref = oriRef
	refType = ref.Type()
	var invalidReflectType reflect.Type

	if ref.Kind() == reflect.Interface {
		ref = ref.Elem()
		if !IsNilReflectValue(ref) {
			refType = ref.Type()
		} else {
			refType = invalidReflectType
		}
	}

	// if its a pointer, resolve its value
	if ref.Kind() == reflect.Ptr {
		// set reftype first. because indirect ref might be nil and we wont be able to get the correct type
		refType = ref.Type().Elem()
		ref = reflect.Indirect(ref)
	}

	if refType != invalidReflectType {
		refKind = refType.Kind()
	} else {
		refKind = reflect.Invalid
	}

	if IsNilReflectValue(ref) {
		refValue = nil
	} else {
		refValue = ref.Interface()
	}

	return ref, refType, refKind, refValue
}

func IsStructReflectValue(ref reflect.Value) bool {
	if ref.Kind() == reflect.Interface {
		ref = ref.Elem()
	}

	// if its a pointer, resolve its value
	if ref.Kind() == reflect.Ptr {
		ref = reflect.Indirect(ref)
	}

	return ref.Kind() == reflect.Struct
}

func GetStructReflectValue(obj interface{}) (ref reflect.Value, err error) {
	ref = reflect.ValueOf(obj)

	if IsNilReflectValue(ref) {
		return ref, errors.New("object cannot be nil")
	}

	if ref.Kind() == reflect.Interface {
		ref = ref.Elem()
	}

	// if its a pointer, resolve its value
	if ref.Kind() == reflect.Ptr {
		ref = reflect.Indirect(ref)
	}

	if ref.Kind() != reflect.Struct {
		return ref, errors.New("object must be of struct type")
	}

	return ref, nil
}
