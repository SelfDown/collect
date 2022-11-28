package test

import (
	"fmt"
	"reflect"
	"testing"
)

type IdCard struct {
	Name string
}
type User struct {
	Id     int
	Name   string
	Weight float32
	Height float32
	IdCard IdCard
}

func Test3(t *testing.T) {
	u1 := &User{
		Id:     7,
		Name:   "zz",
		Weight: 65.5,
		Height: 1.74,
	}
	userMap := make(map[int]*User, 5)
	userMap[u1.Id] = u1
	mapValue := reflect.ValueOf(&userMap)
	card := IdCard{Name: "z"}
	mapValue.Elem().MapIndex(reflect.ValueOf(u1.Id)).Elem().FieldByName("IdCard").Set(reflect.ValueOf(card))
	fmt.Print(userMap)
}
