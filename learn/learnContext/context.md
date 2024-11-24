# context
1. 1个接口: Context
    * context.Context是一个接口
    ```
    type Context interface {
        Deadline()(deadline time.Time(), ok bool)
        Done() <-chan struct{}
        Err() error
        Value(key interface{}) interface{}
    }
    ```
    * 用来在goroutine之间传递消息
        * 为啥要传递消息？不然没法控制子goroutine的运行
2. 4种实现，6个函数（4个具体数据类型，通过6个函数来创建不同功能的Context）
    * emptyCtx
        * type emptyCtx int: 本质上是个int，以下两个函数都会创建emptyCtx
            * func BackGround() Context
                * 没有任何值，截止事件，取消信号
                * 作为顶级任务的根源
            * func TODO() Context
                * 本来应该用外层函数传递的Context，但是外层没有传递
                * 临时占位符
        * 数据结构
        ```
        int
        ```
    * cancelCtx
        * func WithCancel(parent Context)(ctx Context, cancel CancelFunc)
            * 可以把一个Context包装成cancelCtx
            ```
            ctx := context.BackGround()
            ctx1, cancel := context.WithCancel(ctx)
            ```
            * 可以取消
            ```
            case <-ctx.Done() // 接收通道不会继续阻塞
            ```
        * 数据结构
        ```
        type cancelCtx struct {
            Context                           // 父Context
            mu          sync.Mutex            // 保证cancelCtx是线程安全的，随身自带锁，是一个好gopher的习惯
            done        chan stuct{}          // 取消信号
            children    map[canceler]struct{} // 取消根节点时，把子节点都取消掉
            err         error
        }
        ```
    * timerCtx
        * func WithTimeout(parent Context, timeout time.Duration)(Context, CancelFunc)
        * func WithDeadline(parent Context, d time.Time)(Context, CancelFunc)
            * 基于ctx1创建一个timerCtx
            ```
            deadline := time.Now().Add(time.Second)
            ctx2, cancel := WithDeadline(ctx1, deadline)
            ```
        * 数据结构
        ```
        type timerCtx struct {
            cancelCtx
            timer      *time.Timer
            deadline   time.Time
        }
        ```
    * valueCtx
        * 数据结构
        ```
        type valueCtx struct {
            Context
        
            key,val inter
        }

        type iface struct {
            tab * itab
            data
        }

        type itab struct {

        }

        type data struct {
            ctx    // Context
            _type  // key type
            data   // key value
            _type  // value type
            data   // value value
        }
        ```
        * func WithValue(parent Context, key, val interface{}) Context
        ```
        ctxA := WithValue(ctx1, "keyA", "valueA")
        ctxB := WithValue(ctxA, "keyA", "valueB")  // 子节点会覆盖父节点的value值，一般key不用string类型
                                                   // 自定义一个新的类型type keytypec string
        ```
        * Context是本着不可修改来设计的，不要尝试去修改Context中的值
3. 控制goroutine的声明周期
    * 传递取消信号
    ```
    ctx, cancel := context.WithCancel(context.Background())
    go func(){
        for (
            select(
                case <-ctx.Done():
                    goroutine exit
                default:
                    do something
            )
        )
    }()
    cancel()    // 发送取消信号
    ```
    * 传递超时信息
    ```
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
    defer cancel() // 用于清理资源，确保函数返回前执行取消操作
    go func(){
        select{
        case <- time.After(time.Second*1)
            do something
        case <- ctx.Done()
            time out
        }
    }
    ```
    * 传递截止时间
    ```
    deadline := time.Now().Add(time.Second * 5)
    ctx, cancel := context.WithDeadline(context.Background(), deadline)
    defer cancel()
    go func(){
        select {
        case <-time.After(time.Second * 6):
            do something
        case <-ctx.Done():
            deadline exceeded
        }
    }()
    ```
    * 传递数据

4. 监听客户端操作
    ```
	http.ListenAndServe(":8000", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// This prints to STDOUT to show that processing has started
		fmt.Fprint(os.Stdout, "processing request\n")
		// We use `select` to execute a peice of code depending on which
		// channel receives a message first
		select {
		case <-time.After(2 * time.Second):
			// If we receive a message after 2 seconds
			// that means the request has been processed
			// We then write this as the response
			w.Write([]byte("request processed"))
		case <-ctx.Done():
			// If the request gets cancelled, log it
			// to STDERR
			fmt.Fprint(os.Stderr, "request cancelled\n")
		}
	}))
    ```
