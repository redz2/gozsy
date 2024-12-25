# strconv包(字符串的类型转换)
1. string -> []byte or []rune
```
str1 := "Hello World!"
slice1 := []byte(str1)
```

2. Format: other type ---> string
```
// 字符串切片
slice1 := []byte("Hello World")
fmt.Printf("%s\n", slice1)

b := false
b_str := strconv.FormatBool(b)

int_str_basic := strconv.Itoa(123) // 只能转换成十进制
int_str := strconv.FormatInt(140, 16) // 可以十六进制

f_str := strconv.FormatFloat(3.141592, 'f', 4, 64)
```

3. Parse: string ---> other type
```
b1, err := strconv.ParseBool("true") // 如果字符串字面量的值不匹配，返回err
v1, err := strconv.ParseInt("abc", 16, 64) // 十六进制字符串字面量 -> int64
v2, err := strconv.ParseFloat("3.14159", 64) // -> float64
```

4. Append: 将其他类型等转换为字符串后添加到slice
```
slice := make([]byte, 0, 1024)

slice = strconv.AppendBool(slice, false)
slice = strconv.AppendInt(slice, 123, 2)
slice = strconv.AppendFloat(slice, 3.14159, 'f', 4, 64)
slice = strconv.AppendQuote(slice, "hello")
```