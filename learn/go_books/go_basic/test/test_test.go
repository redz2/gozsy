package test

import (
	"testing"
)

// 进入package所在目录执行
// go test github.com/redz2/gozsy/learn/go101/go_basic/test -v

// 单元测试
// go test
// go test -v
// go test -run TestAdd  -v
func TestAdd(t *testing.T) {
	testData := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{4, 5, 9},
		{50, 5, 55},
	}
	ans := 0
	for _, data := range testData {
		if ans = Add(data.a, data.b); ans != data.c {
			t.Errorf("%d + %d expected %d,but %d got", data.a, data.b, data.c, ans)
		}
	}
}

// 基准测试
// 测试命令: go test -bench=. -run=none
// -bench=.: 这个标记表明要进行性能测试（执行该包中所有性能测试函数）
// -run=none: 表示不执行任何功能测试函数（或者这么写: -run=^$）

// goos: darwin
// goarch: arm64
// pkg: github.com/redz2/gozsy/learn/go101/go_basic/test
// cpu: Apple M1 Pro
// BenchmarkAdd-8          1000000000               0.3150 ns/op（执行了这两个测试函数，测试时最大P数量）
// BenchmarkParallel-8     1000000000               0.2253 ns/op（1000000000: 指的是被测函数的执行次数，而不是性能测试函数的执行次数）
// PASS（测试通过）
// ok      github.com/redz2/gozsy/learn/go101/go_basic/test        0.870s（测试总耗时）

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}

// go test -cpu=1,2,4,8 -bench=. -run=none
// 会以1,2,4,8为最大P数量分别去执行第一个测试函数，然后执行第二个函数，以此类推
func BenchmarkParallel(b *testing.B) {
	// 测试一个对象或者函数在多线程的场景下面是否安全
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// m := rand.Intn(100) + 1
			// n := rand.Intn(m)
			_ = Add(1, 2)
		}
	})
}
