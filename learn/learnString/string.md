# 编码规范(字符如何表示)
1. ASCII编码规范
* 单字节字符编码
2. Unicde编码规范
* 实际上是另一个更加通用的、针对书面字符和文本的字符编码标准。它为世界上现存的所有自然语言中的每一个字符，都设定了一个唯一的二进制编码
* 如何表示？U+0061(4个十六进制，共32位)

# 编码格式(字符如何存储)
* 主要关注字符与字节序列之间的转换方式
* UTF-8(以多少位表示一个编码单元，也就是一个字符)
* UTF-8可变宽的编码方案(使用1~4个字节来表示一个字符)

# 关于string一些问题
1. 一个string类型的值在底层是如何表达的？
* 一个string由一系列相对应的Unicode的码点的UTF-8编码值来表示
* []rune: 一个rune是4个字节，可以存放一个Unicode码
* []byte: 一个字节

2. 使用for range时注意
* 先把字符串拆成字符序列，再试图找出这个字节序列中包含的每一个UTF-8的编码值(也及时Unicode字符)

# 类型转换
* string和[]byte可以互相转换
* string和[]rune可以互相转换
    * []byte 和 []rune 不能直接转换，必须通过string
    * []byte <===> string <===> []rune
```
// string to []byte
s1 := "hello"
b := []byte(s1)

// []byte to string
s2 := string(b)
```

# strings包(字符串的处理)
0. 如何使用strings包操作字符串？
    * Count
    * IndexRune
    * Index
        * 在字符串s中查找sep所在的位置，返回位置值，找不到返回-1
        * func Index(s, substr string) int
    * Map
    * Replace
        * 在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
        * func Replace(s, old, new string, n int) string
    * SplitN
    * Split
        * 把s字符串按照sep分割，返回slice
        * func Split(s, sep string) []string
    * Trim
        * 在s字符串的头部和尾部去除cutset指定的字符串
        * func Trim(s string, cutset string) string
    * Contains
        * func Contains(s, substr string) bool
    * Join
        * func Join(a []string, sep string) string
    * Repeat
        * 重复s字符串n次，最后返回重复的字符串
        * func Repeat(s string, count int) string
    * Fields
        * 去除s字符串的空格符，并且按照空格分割返回slice
        * func Fields(s string) []string
        

1. strings.Builder
    * 用来构建字符串，避免频繁的内存分配和拷贝，适用于拼接字符串场景
        * 一旦开始使用builder，不能再复制（如果要重复使用，Reset清空）
    * golang中的字符串，本来就可以做截取(s1[0:3])和拼接(s1+s2)
        * strings.Builder就是来处理大量字符串拼接的场景的
```
func (b *Builder) Grow(n int)
func (b *Builder) Len() int
func (b *Builder) Cap() int
func (b *Builder) Reset()
func (b *Builder) String() string
func (b *Builder) Write(p []byte) (int, error)
func (b *Builder) WriteByte(c byte) error
func (b *Builder) WriteRune(r rune) (int, error)
func (b *Builder) WriteString(s string) (int, error)

var builder strings.Builder  // 创建对象
builder.WriteString("shaw")  // 追加字符串
builder.WriteString("xxx")
fmt.Printf("%s\n", builder.String())
```
2. strings.Reader
    * Builder用来构建字符串
    * Reader用来读取字符串
    * Reader会保存已读取字节的计数(已读计数)
```
func (r *Reader) Len() int // 当前长度
func (r *Reader) Size() int64 // 总长度
func (r *Reader) Read(b []byte) (n int, err error)
func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
func (r *Reader) ReadByte() (byte, error) 
func (r *Reader) UnreadByte() error
func (r *Reader) ReadRune() (ch rune, size int, err error) 
func (r *Reader) UnreadRune() error
func (r *Reader) Seek(offset int64, whence int) 
func (r *Reader) WriteTo(w io.Writer) (n int64, err error)
func (r *Reader) Reset(s string)
func NewReader(s string) *Reader

reader := NewReader("shaw")
b := make([]byte, 3)
n, err := reader.Read(b)
if err != nil {
    fmt.Println("读取失败")
}
fmt.Printf("读取了%d个字节: %s\n", n, string(b)) // 不做格式转换也行


// 通过Size和Len可以计算出已读计数
reader.Size() - int64(reader.Len())
```

# strconv包(字符串的类型转换)
1. 字符串 ---> 字节切片
```
s := "Hello World!"
slice = []byte(s)
```

2. Format: 其他类型转换成字符串
```
// 字符串切片
slice := []byte("Hello World")
fmt.Printf("%s\n", slice)

b := false
b_str := strconv.FormatBool(b)

int_str_basic := strconv.Itoa(123) // 只能转换成十进制
int_str := strconv.FormatInt(140, 16) // 可以十六进制

f_str := strconv.FormatFloat(3.141592, 'f', 4, 64)
```

3. Parse: 将字符串转换成其他类型
```
b1, err := strconv.ParseBool("true")
v1, err := strconv.ParseInt("abc", 16, 64) // 十六进制字符串字面量 -> int64
v2, err := strconv.ParseFloat("3.14159", 64) // -> float64
```

4. Append: 将其他类型等转换为字符串后
```
slice := make([]byte, 0, 1024)

slice = strconv.AppendBool(slice, false)
slice = strconv.AppendInt(slice, 123, 2)
slice = strconv.AppendFloat(slice, 3.14159, 'f', 4, 64)
slice = strconv.AppendQuote(slice, "hello")
```


# 其他开源库
1. go-runewidth
    * 获取字符串宽度
    * 用于cmd展示