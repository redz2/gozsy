# 函数
1. 函数是一等公民
    * 函数基本功能: 封装代码、分割功能、解耦逻辑
    * 作为普通的值: 在函数间传递、赋予变量、做类型判断和转换
    ```
    // 声明一个函数类型
    type Printer func(contents string)(n int, err error)
    
    func printToStd(contents string) (bytesNum int, err error){
        return fmt.Println(contents)
    }

    func main(){
        var p Printer
        // 把函数作为一个普通的值赋给一个变量
        p = printToStd
        p("something")
    }
    ```
2. 什么是高阶函数？ --- 让函数在其他函数之间传递
    * 接受其他的函数作为参数
    * 把其他函数作为结果返回
    ```
    // 对operate的约束
    type operate func(x,y int) int

    func caculate(x,y int, op operate) (int, error) {
        // 卫述语句: 用来检查关键的先决条件的合法性，并在检查未通过的情况下立即终止当前代码的执行
        if op == nil {
            return 0, errors.New("invalid operation")
        }
        return op(x,y), nil
    }
    ```

3. 如何实现闭包？
    * 闭包: 函数中存在对外来标识符的引用(自由变量)
        * 体现从“不确定”变成“确定”的过程
    * 自由变量: 提升变量作用域
    ```
    // genCalculator -> 高阶函数
    // op -> 自由变量
    // calculateFunc -> 闭包函数的类型
    func genCalculator(op operate) calculateFunc {
        return func(x int, y int) (int, error) {
            // 这个函数中的op还不确定
            // 捕获自由变量，形成闭包函数
            if op == nil {
                return 0, errors.New("invalid operation")
            }
            return op(x, y), nil
        }
    }
    ```
    * 实现闭包的意义？
        * 动态生成部分逻辑
        * 模板方法
4. 传入函数的那些参数值后来怎么样了？
    * 所有传给函数的参数值都会被复制(浅层复制)
        * 所以，实际可不可以修改原始数据，和数据本身是值类型还是引用类型有很大关系
        * 如果数组中包含引用类型，引用类型的底层数据也不会拷贝
5. 一个原则
    * 既不要把你程序的细节暴露给外界，也尽量不要让外界的变动影响到你的程序