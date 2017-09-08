package tools

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func OutS(in interface{}) {
	printstruct(reflect.TypeOf(in), reflect.ValueOf(in), 2)
}

func OutA(in interface{}) {
	printarrayslice(reflect.ValueOf(in), 2)
}

func OutM(in interface{}) {
	printmap(reflect.ValueOf(in), 2)
}

func printstruct(t reflect.Type, v reflect.Value, pc int) {
	fmt.Println("")
	for i := 0; i < t.NumField(); i++ {

		fmt.Print(strings.Repeat(" ", pc), t.Field(i).Name+" ("+reflect.TypeOf(v.Field(i).Interface()).String(), "):")
		value := v.Field(i)
		PrintVar(value.Interface(), pc+2)
		fmt.Println("")
	}
}

func printarrayslice(v reflect.Value, pc int) {
	for j := 0; j < v.Len(); j++ {
		PrintVar(v.Index(j).Interface(), pc+2)
	}
}

func printmap(v reflect.Value, pc int) {
	for _, k := range v.MapKeys() {
		PrintVar(k.Interface(), pc)
		PrintVar(v.MapIndex(k).Interface(), pc)
	}
}

func PrintVar(i interface{}, ident int) {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {

		v = reflect.ValueOf(i).Elem()
		t = v.Type()
	}
	switch v.Kind() {
	case reflect.Array:
		printarrayslice(v, ident)
	case reflect.Chan:
		fmt.Println("Chan")
	case reflect.Func:
		fmt.Println("Func")
	case reflect.Interface:
		fmt.Println("Interface")
	case reflect.Map:
		printmap(v, ident)
	case reflect.Slice:
		printarrayslice(v, ident)
	case reflect.Struct:
		if f, ok := i.(time.Time); ok {
			fmt.Print(strings.Repeat(" ", ident), f.Format("2006-01-02 15:04:05"))
		} else {
			printstruct(t, v, ident)
		}
	case reflect.UnsafePointer:
		fmt.Println("UnsafePointer")
	default:
		fmt.Print(strings.Repeat(" ", ident), v.Interface())
	}
}
