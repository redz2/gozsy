# 容器
1. 容器类型和容器值概述
    * Go中的容器类型: 数组、切片、映射
    * 每个容器用来表示和存储一个元素序列或集合
        * 每个元素值都关联一个key（数组和切片使用int，映射必须使用可比较的key）
        * 每个容器值都有一个长度属性
        * 内存结构上的差别
            * 数组的所有元素都在直接部分（这就是为啥说数组是值类型，而切片是引用类型）
            * 切片和映射的元素都在间接部分

2. 无名容器类型的字面表示形式
    * [N]T
    * []T
    * map[K]T
    * 所有切片类型的尺寸都是一致的，所有映射类型的尺寸都是一致的，数组类型的尺寸等于元素长度*元素个数
        * 何为尺寸？unsafe.Sizeof()
    ```
    const Size = 32

    type Person struct {
        name string
        age  int
    }

    // 数组类型
    [32]string
    [Size]int
    [16][]byte
    [100]Person

    // 切片类型
    []bool
    []int64
    []map[int]bool
    []*int

    // 映射类型
    map[string]int
    map[int]bool
    map[int16][6]string
    map[struct {x int}]*int8
    ```

3. 容器字面量的表示形式
    * T{...} 大括号中的每一项称为一个键值元素对
    ```
    [4]bool{true, true, true, true}
    []string{"xxx", "yyy"}
    map[string]int{"xxx": 1, "yyy": 2}

    [...]bool{true, true} // 这也是数组，编译器会计算出长度
    ```

4. 容器类型零值的字面量表示形式
    * 数组类型A的零值: A{}
    * 切片和映射类型的零值: nil
    * []T{} 和 []T(nil) 
        * 空切片值和nil是不等价的: nil切片标识(间接部分)内存空间尚未开辟
        ```
        []int{} == nil // false
        ```

5. 容器的字面量是不可寻址的但可以被取地址
    * 字面量只是表示形式，哪来的地址啊，下面只是一种写法而已（变量才有地址）
    ```
    pm := &map[string]int{"xxx": 1}  // var tmp := map[string]int{"xxx": 1}; pm := &tmp
    ```

6. 内嵌组合字面量可以被简化
    ```
    var heads = []*[]byte{
        &[4]byte{'P', 'N', 'G', ''},
        &[4]byte{'G', 'I', 'F', ''},
        &[4]byte{'J', 'P', 'E', 'G'},
    }
    // 类型部分可以省略掉！！！这样就很方便
    var heads = []*[]byte{
        {'P', 'N', 'G', ''},
        {'G', 'I', 'F', ''},
        {'J', 'P', 'E', 'G'}, // 最后一行的逗号不能省略
    }
    ```

7. 容器值的比较
    * slice和map无法比较，只能和nil进行比较

8. 查看容器值的长度和容量
    * len: 包含多少个元素
    * cap: 容器的容量

9. 读取和修改容器的元素

10. 什么情况下切片为nil
    ```
    var a = []int{}  // 不是nil
	var b []int      // nil
    ```
