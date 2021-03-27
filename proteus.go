package proteus

import (
	"reflect"
)

type Proteus interface {
	Map(interface{}, interface{})
}

type proteanImpl struct {
	tagName string
}

func (p proteanImpl) mapValue(src interface{}, dst reflect.Value) {
	getValue := func(val reflect.Value) reflect.Value {
		if val.Kind() == reflect.Ptr {
			return reflect.Indirect(val)
		}
		return val
	}
	var srcValue = getValue(reflect.ValueOf(src))
	numFields := srcValue.NumField()
	srcType := srcValue.Type()
	for i := 0; i < numFields; i++ {
		field := srcType.Field(i)
		if len(field.PkgPath) != 0 {
			continue
		}
		if tag, exists := field.Tag.Lookup(p.tagName); exists {
			value := getValue(srcValue.Field(i))
			if len(tag) > 0 {
				targetField := dst.FieldByName(tag)
				if !targetField.IsValid() || !targetField.CanSet() {
					continue
				}
				targetType := targetField.Type()
				if targetType.AssignableTo(value.Type()) {
					targetField.Set(value)
				} else if value.Type().Kind() == reflect.Struct && (targetType.Kind() == reflect.Struct || targetType.Kind() == reflect.Ptr) {
					p.mapValue(value.Interface(), getValue(targetField))
				}
			} else if field.Type.Kind() == reflect.Struct {
				p.mapValue(value.Interface(), dst)
			}
		}
	}
}

func (p proteanImpl) Map(src interface{}, dst interface{}) {
	val := reflect.ValueOf(dst)
	if val.Kind() == reflect.Ptr {
		p.mapValue(src, val.Elem())
	}
	// No need to do anything since dst isn't a pointer anyway
}

func New(tag string) Proteus {
	return proteanImpl{tagName: tag}
}
