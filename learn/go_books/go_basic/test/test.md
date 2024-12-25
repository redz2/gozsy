# 测试
1. 功能测试（test）
    * 对于功能测试函数来说，其名称必须以Test为前缀，并且参数列表中只应有一个*testing.T类型的参数声明
2. 基准测试（benchmark，也称为性能测试）
    * 对于性能测试函数来说，其名称必须以Benchmark为前缀，并且唯一参数的类型必须是*testing.B类型的
3. 示例测试（example）
    * 对于示例测试函数来说，其名称必须以Example为前缀，但对函数的参数列表没有强制规定