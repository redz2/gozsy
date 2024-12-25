# go-web
1. Registering a Request Handler: 处理请求
2. Listen for HTTP Connections: 服务监听
3. gorilla/mux: 更高级的路由匹配
    * httprouter
4. go get -u github.com/go-sql-driver/mysql: 数据库操作
5. html/template: 模板
6. Assets and Files: 静态文件
7. middleware: 中间件
    * 装饰器思想
8. Sessions: gorilla/sessions 
9. websocket: github.com/gorilla/websocket
10. Password Hashing (bcrypt): golang.org/x/crypto/bcrypt

# web框架分层
* protocol -> middleware -> controller -> logic -> dao -> storage
```
// 如何在入口支持多种协议？grpc、http
// protocol: 把数据从协议特定的结构体读出来，写入到和协议无关的结构体上
func HTTPCreateOrderHandler(wr http.ResponseWriter, r *http.Request) {
    var req CreateOrderRequest
    var params CreateOrderParams
    ctx := context.TODO()
    // bind data to req
    bind(r, &req)
    // map protocol binded to protocol-independent
    map(req, params)
    logicResp,err := controller.CreateOrder(ctx, &params)
    if err != nil {}
    // ...
}

// middleware: 中间件层（非功能性逻辑代码）

// controller: 和协议无关
func CreateOrder(ctx context.Context, req *CreateOrderStruct) (
    *CreateOrderRespStruct, error,
) {
    // ...
}
```

* 业务系统发展
    * 业务无关: 拆解和异步化（比如统计、用户状态更新）
        * 时延敏感: rpc
        * 不敏感: 消息队列
    * 单个业务变得复杂时，使用函数封装业务流程（无法继续拆分了咋办？）
        * step1 -> step2 -> step3
        * 可能会发现多个业务具有相似的step
    * 使用接口来做抽象
        * 千万不要过早地引入接口
        * 不同业务可能有相似的流程
            * 部分流程不相似，强制要求实现接口中的函数，可以为空
        ```
        // 业务流程的抽象
        type BusinessInstance interface {
            ValidateLogin()
            ValidateParams()
            AntispamCheck()
            GetPrice()
            CreateOrder()
            UpdateUserStatus()
            NotifyDownstreamSystems()
        }
        // 在业务入口处判断业务类型
        func entry() {
            var bi BusinessInstance
            switch businessType {
                case TravelBusiness:
                    bi = travelorder.New()
                case MarketBusiness:
                    bi = marketorder.New()
                default:
                    return errors.New("not supported business")
            }
        }
        // 业务实现（不然的话，每个业务流程都需要判断是什么具体类型？）
        // 如果没有接口，那么
        func BusinessProcess(bi BusinessInstance) {
            bi.ValidateLogin()
            bi.ValidateParams()
            bi.AntispamCheck()
            bi.GetPrice()
            bi.CreateOrder()
            bi.UpdateUserStatus()
            bi.NotifyDownstreamSystems()
        }
        ```

* 表驱动开发
    * 如何干掉函数中的if和switch？
    ```
    func entry() {
        var bi BusinessInstance
        switch businessType {
        case TravelBusiness:
            bi = travelorder.New()
        case MarketBusiness:
            bi = marketorder.New()
        default:
            return errors.New("not supported business")
        }
    }

    // 使用表来取代if和switch
    var businessInstanceMap = map[int]BusinessInstance {
        TravelBusiness : travelorder.New(),
        MarketBusiness : marketorder.New(),
    }

    func entry() {
        bi := businessInstanceMap[businessType]
    }
    ```
    
