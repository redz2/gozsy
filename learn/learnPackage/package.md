# package
1. 有哪些类型的包？
    * 系统内置
    * 自定义包
    * 第三方包
        * go get
        * go mod download
        * go mod vendor
2. 初始化一个项目
    * go mod init path/module
        * 单个模块: go mod init github.com/{your-username}/{repo-name}
        * 多个模块: go mod init github.com/{your-username}/{repo-name}/{module-name}
    * 目录结构
        * 同一个目录下只能有一个package（一般和目录同名）
            * package.go
            * package_test.go
            * main.go（比较特殊）
        * 同一个项目中创建多个app
            * 一般在cmd目录中创建多个目录，每个目录对应一个app
    ```
    module
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
3. main.go必须在module目录下，否则无法导入自定义包
4. github.com/redz2/learn-python
    * github地址 + 用户 + 项目名
    * go mod init github.com/redz2/gozsy