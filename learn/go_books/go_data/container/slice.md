# 数组、字符串、切片
1. 数组、字符串、切片这三种数据类型在底层原始数据有着相同的内存结构
    * 数组: 是字符串和切片的基础（底层间接值部）
    * 字符串: 只读，不能修改
    * 切片: 结构和字符串差不多
2. 数组
    * 数组的定义
    ```
    var a [3]int
    var b = [...]int{1, 2, 3}
    var c = [...]int{2: 3, 1: 2}
    var d = [...]int{1, 2, 4: 5, 6}

    var a_ptr = &a  // 数组指针并不是数组，不过操作方式和数组基本一致
    for i := range a_ptr {
        fmt.Println(i)
    }
    ```
3. 字符串
    * 一个字符串是一个不可改变的字节序列（一个只读的字节数组）
    ```
    var data = [...]byte{
        'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd',  // 和字符串底层数据完全一致
    }
    ```
    * string、[]byte、[]rune
        * 字节和字符的区别？
        * 如何识别一个字符？UTF-8针对字符有一套模板
4. 切片
    * 切片的定义
    ```
    var (
        a []int          // nil切片
        b = []int{}      // 空切片，和nil切片不同（是否存在底层数组的区别）
        c = []int{1,2,3}
        d = c[:2]        // 切片底层共享数组，如果多个切片很容易造成数据错误
        e = c[0:2:cap(c)]
        f = c[:0]
        g = make([]int,3)
        h = make([]int,2,3)
        i = make([]int,0,3)
    )
    // 添加切片元素
    a = append(a, 1, 2)         // append会创建底层数组
    b = append([]int{0}, b...)  // 在b的开头添加一个item
    // 删除切片元素
    g = g[:len(g)-1]            // 删除最后一个item（底层数组并没有删除任何东西）
    g = g[1:]                   // 删除最后一个item
    ```
    * 一个删除[]byte空格的函数
        * 使用切片需要关注当前使用的底层数组是否还有其他切片在使用（如果没有无须过多担心）
    ```
    func TrimSpace(s []byte) []byte {
        // 共享底层数组，避免分配内存（共用底层数组）
        b := s[:0]
        for _, x := range s {
            if x != ' ' {
                b = append(b, x)
            }
        }
        return b
    }

    func main() {
        s := []byte("xx xxxsadsa  yy")
        x := TrimSpace(s)
        fmt.Println(string(x))
        fmt.Println(string(s))   // s的底层数组已经被修改的面目全非了，感觉很容易出bug
    }
    ```