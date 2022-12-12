package test

import (
	"fmt"
	"reflect"
	"testing"
)

func test() int {

	return 1
}
func TestT(t *testing.T) {
	i := test()
	t2 := &i
	var t3 int32
	t3 = 12
	t4 := &t3
	fmt.Println(i)
	fmt.Printf("%#v\n", t2)
	fmt.Printf("%#v\n", t4)
	reflect.ValueOf(t2)
	tt := reflect.New(reflect.TypeOf(t2)).Elem().Type()
	reflect.New(tt).Elem().Convert(reflect.TypeOf(t3))
}
