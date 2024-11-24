# Unintended variable shadowing（变量遮蔽）
1. 外层的client永远是nil
    ```
	var client *http.Client
	if tracing {
		client, err := createClientWithTracing()
		if err != nil {
			return err
		}
		log.Println(client)
	} else {
		client, err := createDefaultClient()
		if err != nil {
			return err
		}
		log.Println(client)
	}
	// Use client
    ```
2. 如何解决？
    ```
    var client *http.Client
    if tracing {
        c, err := createClientWithTracing()
        if err != nil {
            return err
        }
        client = c  // 创建一个临时变量
    }
    ```

    ```
    var client *http.Client
    var err error
    if tracing {
        client, err = createClientWithTracing()    // 使用赋值，而不是短变量声明
        if err != nil {
            return err
        }
    }

    // 更简单的方式，只处理一次错误
    var client *http.Client
    var err error
    if tracing {
        client, err = createClientWithTracing()    // 使用赋值，而不是短变量声明
    } else {
        client, err = createDefaultClient()
    }
    if err != nil {
        return err
    }
    ```