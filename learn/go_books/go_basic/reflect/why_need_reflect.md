# why we need reflect
1. 使用switch类型分支来判断类型，但我们没有办法来检查未知类型
```
type stringer interface {
    String() string
}

func Sprint(x interface{}) string {
    switch x := x.(type) {
    case stringer:
        return x.String()
    case string:
        return x
    case int:
        return strconv.Itoa(x)
    // ...similar cases for int16, uint32, and so on...
    case bool:
        if x {
            return "true"
        }
        return "false"
    default:
        // array, chan, func, map, pointer, slice, struct
        return "???"
    }
}
```

2. reflect.Type 和 reflect.Value
    * 都满足 fmt.Stringer 接口
    * fmt.Printf 提供了一个缩写 %T 参数, 内部使用 reflect.TypeOf 来输出
    ```
    fmt.Printf("%T\n", 3)   // "int"  reflect.Type
    fmt.Printf("%v\n", v)   // "3"    reflect.Value
    ```