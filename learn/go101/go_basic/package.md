# package
1. package类型
    * builtin
    * 第三方package
    * 自定义package

2. 初始化一个项目
    * go mod init path/to/module
        * 单个模块: go mod init github.com/{your-username}/{module-name}
    * go mod init path/to/repo
        * 多个模块: go mod init github.com/{your-username}/{repo-name}/{module-name}
        ```
        repo-name
        - module1
          - go.mod
          - go.sum
          - module1.go
        - module2
          - go.mod
          - go.sum
          - module2.go
        ```
    * 目录结构
        * 一个目录就是一个package（package一般和目录同名，一个目录下的文件必须都是同一个package）
            * package.go
            * package_test.go
            ```
            myapp
            - cmd
              - app01
                - main.go
              - app02
                - main.go
            - package
              - package1.go
                package2.go
            - api
              - api1.go
                api2.go
            main.go
            go.mod
            go.sum
            ```
        * main.go如何导入本地package？
          * go.mod: 所在目录就是ROOT目录，import模块时会从根目录开始找
          * import: 模块名/相对目录
        * 项目和模块的区别？
          * 一个项目包含main.go文件，能够编译成可执行文件
          * 一个模块通常只包含package文件

3. 自己创建一个模块上传至github
    * 在一个项目中创建多个模块
      * 需要创建多个目录，每个目录进行go mod init
    ```
    cd learn/hello/
    go mod init github.com/redz2/gozsy/learn/hello

    cd learn/world/
    go mod init github.com/redz2/gozsy/learn/world
    ```
    * 提交代码至github
    * 配置tag（这一步比较关键）
    ```
    git tag learn/hello/v1.0.0
    ```
    * 发布版本（release）
    * 如何使用发布的模块
    ```
    go get github.com/redz2/gozsy/learn/hello@v1.0.0
    ```
    * go mod vendor: 将模块依赖的第三方库下载到vendor目录下(离线开发)
    ```
    go mod vendor  // make vendored copy of dependencies
    go mod tidy // add missing and remove unused modules
    go mod download
    ```

4. 代码包和包引入
  * 包引入
    * 完整的引入声明语句
    ```
    import importname "path/to/package"  // 默认值是引入包的包名，一般省略
    ```
  * init函数
  * 每个非匿名引入必须至少使用一次