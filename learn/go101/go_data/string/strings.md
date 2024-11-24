# strings包(字符串的处理)
1. 如何使用strings包操作字符串？
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

2. strings.Builder
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
3. strings.Reader
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



# 其他开源库
1. go-runewidth
    * 获取字符串宽度
    * 用于cmd展示