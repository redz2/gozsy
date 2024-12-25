# sort.Interface接口
1. 一个内置的排序算法需要知道: 序列的长度，表示两个元素比较的结果，一种交换两个元素的方式
```
type Interface interface {
    Len() int
    Less(i, j int) bool // i, j are indices of sequence elements
    Swap(i, j int)
}

type StringSlice []string
func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
sort.Sort(StringSlice(names))

sort.Strings(names) // 因为经常使用，已经提供了一个方法

sort.Ints([]int{1,2,3})
sort.IntSlice([]int{1,2,3})
sort.Reverse(sort.Ints([]int{1,2,3}))
```

2. 当实现sort.Interface接口就能使用排序算法
    * 说明针对于该接口实现了排序函数
    * 为何内置？因为用的多