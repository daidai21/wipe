package wipe

import (
	"errors"
	"reflect"
	"time"
)

var (
	NotKnowTypeFieldErr = errors.New("not know type field err")
)

// Wipe val必须是结构体的地址
func Wipe(val interface{}) error {
	v := reflect.ValueOf(val).Elem()

	// fixme: time.Time 这种常用的单独处理下
	if v.Kind() == reflect.Struct {
		if _, ok := val.(*time.Time); ok {
			v.Set(reflect.Zero(v.Type()))
		}
	}

	return wipeAtStruct(v)
}

func wipeAtInterface(val interface{}) error {
	// fixme: error接口置空
	return nil
}

func wipeAtStruct(v reflect.Value) error {
	var fieldErr error

	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		if !fv.CanSet() {
			continue
		}
		if fv.IsZero() {
			continue
		}
		switch fv.Kind() {
		case reflect.Bool:
			fv.Set(reflect.ValueOf(false))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fv.SetInt(0)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			fv.SetUint(0)
		case reflect.Uintptr:
			fv.Set(reflect.ValueOf(nil))
		case reflect.Float32, reflect.Float64:
			fv.Set(reflect.ValueOf(0.0))
		case reflect.Complex64, reflect.Complex128:
			fv.Set(reflect.ValueOf(complex(0, 0)))
		case reflect.Array, reflect.Slice:
			fv.SetLen(0)
			fv.SetCap(0)
		case reflect.Chan:
			fv.Close()
		case reflect.Func:
			fv.Set(reflect.ValueOf(func() {}))
		case reflect.Interface:
			fieldErr = wipeAtInterface(fv)
		case reflect.Map:
			fv.Set(reflect.MakeMapWithSize(fv.Type(), 0))
		case reflect.Ptr:
			fv.Elem().Set(reflect.Zero(fv.Elem().Type()))
			//fieldErr = wipeAtPtr(fv)
		case reflect.String:
			fv.SetString("")
		case reflect.Struct:
			fieldErr = wipeAtStruct(fv)
		case reflect.UnsafePointer:
			fv.Set(reflect.ValueOf(nil))
		default:
			return NotKnowTypeFieldErr
		}

		if fieldErr != nil {
			return fieldErr
		}
	}
	return nil
}

func wipeAtPtr(v reflect.Value) error {
	ve := v.Elem()
	if ve.Kind() == reflect.Struct {
		return wipeAtStruct(ve)
	} else {
		ve.Set(reflect.Zero(ve.Type()))
		//newFv := reflect.New(v.Type().Elem())
		//v.Set(newFv)
		return nil
	}
}
