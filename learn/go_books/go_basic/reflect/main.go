package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
	age  int
}

type manager struct {
	user
	titile string
}

func (m manager) String() string {
	return "ok"
}

func main() {
	var m manager = manager{
		user{
			"xxx",
			13,
		},
		"book",
	}
	// 返回变量的类型 -> manager
	t := reflect.TypeOf(&m)
	// 检查是不是指针
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// 检查manager的字段
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		// 每个字段都有名称，类型
		fmt.Println(f.Name, f.Type, f.Offset)
		// 如果是匿名字段
		if f.Anonymous {
			// f.Type -> user
			for x := 0; x < f.Type.NumField(); x++ {
				// user的每个字段也有名称，类型
				af := f.Type.Field(x)
				fmt.Println(af.Name, af.Type, af.Offset)
			}
		}
	}

	// 反射可以探知外包的非导出结构成员
	// 输出方法集
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Println(t.Method(i))
	}

	// 如何通过反射修改int类型的值？
	num := 100
	reflect_num := reflect.ValueOf(&num).Elem()
	reflect_num.SetInt(200)
	fmt.Println(num)

	// 如何通过反射修改slice的值？
	my_slice := []int{1, 2, 3, 4, 5}
	reflect_slice := reflect.ValueOf(my_slice) // slice是引用类型，那我修改reflect.Value中的值，就是修改了本身的值
	fmt.Println(reflect_slice.Slice(1, 3))
	e := reflect_slice.Index(0) // 如何修改reflect.Value的值？找到索引
	e.SetInt(20000)
	fmt.Println(my_slice)

}
