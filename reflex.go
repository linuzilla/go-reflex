package reflex

import (
	"log"
	"reflect"
)

type Reflex interface {
	Instance() interface{}
	NewInstance() interface{}
	Append(instance interface{})
}
type reflexImpl struct {
	instance       interface{}
	pointer        reflect.Value
	reflectValue   reflect.Value
	reflexIndirect reflect.Value
	dataType       reflect.Type
}

func New(data interface{}) Reflex {
	reflectValue := reflect.ValueOf(data)

	if reflectValue.Kind() != reflect.Ptr {
		log.Fatal(reflectValue.Kind().String() + " not support")
	}

	reflectIndirect := reflect.Indirect(reflectValue)

	switch reflectIndirect.Kind() {
	case reflect.Array, reflect.Slice:
		dataType := reflect.TypeOf(data).Elem().Elem()
		pointer := reflect.New(dataType)

		return &reflexImpl{
			instance:       pointer.Interface(),
			pointer:        pointer,
			reflectValue:   reflectValue,
			reflexIndirect: reflectIndirect,
			dataType:       dataType,
		}
	case reflect.Struct:
		dataType := reflect.TypeOf(data).Elem()
		pointer := reflect.New(dataType)

		return &reflexImpl{
			instance:       pointer.Interface(),
			pointer:        pointer,
			reflectValue:   reflectValue,
			reflexIndirect: reflectIndirect,
			dataType:       dataType,
		}
	default:
		log.Fatal(reflectIndirect.Kind().String() + " not support")
	}

	return nil
}

func (ref *reflexImpl) Instance() interface{} {
	return ref.instance
}

func (ref *reflexImpl) NewInstance() interface{} {
	return ref.pointer.Interface()
}

func (ref *reflexImpl) Append(instance interface{}) {
	ref.reflexIndirect = reflect.Append(ref.reflexIndirect, reflect.Indirect(reflect.ValueOf(instance)))
	reflect.Indirect(ref.reflectValue).Set(ref.reflexIndirect)
}
