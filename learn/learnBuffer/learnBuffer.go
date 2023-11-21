package main

import (
	"bytes"
	"fmt"
	"os"
)

// type Buffer struct {
// 	buf      []byte // contents are the bytes buf[off : len(buf)]
// 	off      int    // read at &buf[off], write at &buf[len(buf)]
// 	lastRead readOp // last read operation, so that Unread* can work correctly.
// }

// 缓冲区的零值是一个待使用的空缓冲区
// 从bytes.Buffer读取数据后，被成功读取的数据仍保留在原缓冲区，只是无法被使用，因为缓冲区的可见数据从偏移off开始，即buf[off : len(buf)]。

func main() {

	// 声明一个Buffer
	// var b bytes.Buffer
	// b := new(bytes.Buffer)
	// b := bytes.NewBuffer([]byte{'z', 'h', 'o', 'i'})
	b := bytes.NewBufferString("zhouyi")

	// 往buffer中写入数据
	// 将切片写入Buffer尾部
	// b.Write([]byte{' ', 'o', 'k'})
	// 将字符串写入Buffer尾部
	b.WriteString(" ok")
	// 将字符写入Buffer尾部
	// b.WriteByte('B')
	// b.WriteByte(66)
	// 将rune类型的数据写入Buffer尾部
	// b.WriteRune(400)
	// 从io.Reader接口读取数据写入Buffer尾部
	file, _ := os.Open("./test.txt")
	b.ReadFrom(file)

	// 从Buffer中读取数据
	b.ReadByte()
	// 读取n个字节并返回
	// b.Next(n int) []byte
	// 一次性读取len(p)数据到p中
	// b.Read(p []byte) (n int, err error)

	// 读取一个Unicode字符并返回
	// b.ReadRune() (r rune, size int, err error)
	// 读取一个字节并返回
	// b.ReadByte() (byte, error)
	// 读取缓冲区第一个分隔符前面的内容以及分隔符并返回，缓冲区会清空读取的内容。如果没有发现分隔符，则返回读取的内容并返回错误io.EOF
	// b.ReadBytes(delimiter byte) (line []byte, err error)
	// 读取缓冲区第一个分隔符前面的内容以及分隔符并作为字符串返回，缓冲区会清空读取的内容。如果没有发现分隔符，则返回读取的内容并返回错误 io.EOF
	// b.ReadString(delimiter byte) (line string, err error)

	//将Buffer中的内容输出到实现了io.Writer接口的可写入对象中，成功返回写入的字节数，失败返回错误
	// b.WriteTo(w io.Writer) (n int64, err error)

	fmt.Printf("first b: %v\n", b)

	// 其他操作
	// 返回字节切片
	fmt.Printf("b.Bytes(): %v\n", b.Bytes())
	// 返回Buffer内部字节切片的容量
	fmt.Printf("b.Cap(): %v\n", b.Cap())
	// 一般不会自己增加，不会用到
	b.Grow(33)
	fmt.Printf("b.Cap(): %v\n", b.Cap())
	// 返回Buffer内部字节切片的长度
	fmt.Printf("b.Len(): %v\n", b.Len())
	// 清空数据
	// b.Reset()
	// b.String()
	// b.UnreadByte()
	b.UnreadByte()
	// b.UnreadRune()
	fmt.Printf("second b: %v\n", b)

}
