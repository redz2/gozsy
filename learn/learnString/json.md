# json
1. 序列化 struct -> []byte
    * 如果tag中写omitempty，序列化成字符串会忽略该字段
    * struct中一定是完整的，json字符串中可以省略

2. 反序列化 []byte -> struct
    * 相反，如果字符串没有该字段，反序列化时会补充为默认零值，没有omitempty则必须要写全

3. Example
    ```
    type Person struct {
        Name   string `json:"name"`
        Age    string `json:"age"`
        Weight string `json:"weight"`
    }

    func transformation() {
        data_struct := Person{
            Name:   "bob",
            Age:    "23",
            Weight: "160",
        }
        //  将结构体转换成为json
        data_byte, _ := json.Marshal(data_struct)
        fmt.Printf("str:%v\n", string(data_byte))
        
        var p2 Person
        json.Unmarshal(data_byte, &p2)
        fmt.Printf("struct:%v\n", p2)
    }
    ```

4. 结构体嵌套
    * 匿名无tag，字符串是单层结构
    * 匿名有tag，字符串是嵌套结构（一般会用这个，能显示层级结构）

5. tag中指定响应的数据类型(类型不一致)
    ```
    type Card struct {
        ID    int64   `json:"id,string"`
        Score float64 `json:"score,string"`
    }

    `{"id":"233","score":"100"}`
    ```