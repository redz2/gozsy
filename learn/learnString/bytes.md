# bytes包
1. strings vs bytes
    * strings包面向Unicode字符和UTF-8编码的字符串
        * 读: strings.Reader
        * 写: strings.Builder
    * bytes包面向字节和字节切片
        * 集读写功能于一身的数据类型
2. bytes.Buffer
    * 用作字节序列的缓冲区
        * []byte
    * bytes.Buffer vs strings.Builder
        * strings.Builder只能拼接和导出字符串
        * bytes.Buffer不但可以拼接、截断其中的字节序列，以各种形式导出其中的内容，还可以顺序地读取其中的子序列
        * bytes.Buffer无法轻易获取已读计数(但是内部会保存已读计数，很多方法都基于此数)
```
var buffer1 bytes.Buffer                                   // 声明一个buffer
contents := "Simple byte buffer for marshaling data." 
fmt.Printf("Writing contents %q ...\n", contents)
buffer1.WriteString(contents)                              // 写入一个字符串
fmt.Printf("The length of buffer: %d\n", buffer1.Len())    // 长度: 内容容器中未读内容 39
fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())  // 容量: 内容容器的容量，不是内容的长度(无法获取到内容容量) 64


p1 := make([]byte, 7)
n, _ := buffer1.Read(p1)                                   // 读取7个字节
fmt.Printf("%d bytes were read. (call Read)\n", n)
fmt.Printf("The length of buffer: %d\n", buffer1.Len())    // 长度: 32
fmt.Printf("The capacity of buffer: %d\n", buffer1.Cap())  // 容量: 64
```
3. 已读计数作用
    * 读取内容: 根据已读计数找到未读内容，并更新计数
        * Read
        * Next
        * WriteTo
    * 写入内容: 检查内容容器容量是否足够，扩容时只会把未读部分拷贝到新的容器中，已读计数设置为0
        * Write
        * RreadFrom
    ```
    var buffer bytes.Buffer
    buffer.readFrom(strings.NewReader("this is string"))
    ```
    * 截断内容: 截断已读计数之后未读部分 
        * Truncate
    * 读回退
    * 重置内容: 已读计数设置为0
    * 导出内容: 导出已读计数之后未读部分
    * 获取长度: 根据已读计数和内容的长度，计算出未读部分长度
    ```
    操作的都是未读部分的数据(u的那部分)

    [ddddduuuuuuuuuuu00000000000000000]

    ```

4. bytes.Buffer扩容策略(底层会做许多事情)
* 如果内容容器的容量与长度的差小于所需的字节数
```
buf = buf[:length+need] // 数组已经够大的情况，如果不够就会扩容
                        // 优化点: 如果容器容量的一半大于未读容量+所需字节数，扩容代码会使用容器已读容量部分，覆盖调之前的内容
                        // 零值buffer，默认会初始化一个64字节的容器容量
```

5. bytes.Buffer常用方法
* bytes.Buffer实现了io.Reader和io.Writer接口
* 读取
```
buf := bytes.NewBufferString("shaw")
fmt.Printf("%s\n", buf.String())

s := new([]byte, 4)
n, _ := buf.Read(s)

// buf.Next(2)
b, _ := buf.ReadByte()
r, _ := buf.ReadRune() // 读取一个Rune，返回大小(应该是用来计数的，底层是UTF-8编码的字节序列)
// buf.UnReadRune()    // 回退一个Rune
str, _ := buf.ReadString('\n')

```
* 写入
```
var buf bytes.Buffer
buf.Write([]byte("hello "))

buf.WriteByte('w')
buf.WriteString("orld!")
buf.WriteRune([]rune("shaw"))

// 最通用的方法
// func (b *Buffer) ReadFrom(r io.Reader) (n int64, err error)
```