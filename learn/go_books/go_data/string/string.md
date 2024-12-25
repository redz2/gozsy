# string
* 字符串的字面量
    * 双引号（解释型）
    * 单引号（所见即所得）
    ```
    s1 := "zzz\n"
    s2 := `yyy\n`
    ```
* 字符串的零值: "" or ``，不是nil
* 拼接字符串: "a" + "b"
```
s1 := "hello"
s2 := "hello"
fmt.Sprintf("%s %s", s1, s2)
strings.Join(s1, s2)

builder := strings.Builder
builder.WriteString(s1)
builder.WriteString(s2)

buffer := bytes.Buffer
buffer.WriteString(s1)
buffer.WriteString(s2)
```
* 字符串内容（底层字节）不可变
    * 其实是一种约定
    * string的实现不包含内存空间，只有内存指针（指向字符串字面量，而字符串字面量存储在只读段）
* string没有内置方法
    * strings标准库
    * len获取字节数
    * s[1]获取字节
* 字符串转换
    * []rune: 4字节，一个Unicode码点（非法码点会认为是乱码）
        * 读取字符串中的每个字符，将每个字符补充为4字节
    * []byte
        * 需要一次内存拷贝，根本原因在于string不可变，需要为转换后的类型分配内存
    * []rune和[]byte不能直接转换: []byte <===> string <===> []rune
    ```
    // 只能把[]byte -> []rune
    rs := bytes.Rune([]byte("xxx"))

    // utf-8包，把[]rune -> []byte
    func Runes2Bytes(rs []rune) []byte {
        n := 0
        for _, r := range rs {
            n += utf8.RuneLen(r)  // 获取每个码点的字节数
        }
        n, bs := 0, make([]byte, n)  // 创建一个对应长度的切片
        for _, r := range rs {       // 每个rune
            n += utf8.EncodeRune(bs[n:], r)  // 把每个rune转换成字节添加到对应的位置
        }
        return bs
    }
    ```

