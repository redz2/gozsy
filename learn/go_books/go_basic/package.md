# package
1. 包有哪些？
    * Go内置
    * 开源社区
    * 自己写的

2. 初始化一个项目
    * go mod init path/to/module
        * go mod init github.com/{your-username}/{module-name}
        ```
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
        * 构建二进制可执行文件的go项目结构
        ```
        GoProj
        cmd/
        - app01/
          - main.go
        - app02/
          - main.go
        pkg/
        - lib1/
          - lib1.go
        - lib2/
          - lib2.go
        go.mod
        go.sum
        ```
        * 构建库的Go项目结构
        ```
        lib.go
        lib1/
        - lib1.go
        lib2/
        - lib2.go
        go.mod
        go.sum
        ```
        * main.go如何导入本地package？
          * go.mod: 所在目录就是ROOT目录，import模块时会从根目录开始找
          * import: 模块名/相对目录（需要注意，不是绝对路径，模块名称+相对路径）

3. 自己创建一个模块上传至github
    * 一个项目中包含多个模块
      * 创建多个目录，每个目录进行go mod init
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

5. internal目录
  * 不想暴露给外部引用